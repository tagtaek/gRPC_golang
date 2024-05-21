package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"project4/bidirectional"

	"google.golang.org/grpc"
)

func makeMessage(message string) *bidirectional.Message {
	return &bidirectional.Message{
		Message: message,
	}
}

func generateMessages() chan *bidirectional.Message {
	messages := make(chan *bidirectional.Message)

	go func() {
		defer close(messages)

		messagesToSend := []string{"message #1", "message #2", "message #3", "message #4", "message #5"}
		for _, msg := range messagesToSend {
			messages <- makeMessage(msg)
			fmt.Printf("[client to server] %s\n", msg)
		}
	}()

	return messages
}

func sendMessage(client bidirectional.BidirectionalClient) {
    ctx := context.Background()

    stream, err := client.GetServerResponse(ctx)
    if err != nil {
        log.Fatalf("스트림을 열 수 없습니다.: %v", err)
    }

    // 메시지 채널을 통해 메시지를 서버로 전송합니다.
    for msg := range generateMessages() {
        if err := stream.Send(msg); err != nil {
            log.Fatalf("메시지를 보낼 수 없습니다.: %v", err)
        }
    }

	// -------- client가 자동 종료되도록 추가한 코드 --------
	// 클라이언트에서 스트림을 닫습니다.
	if err := stream.CloseSend(); err != nil {
		log.Fatalf("스트림을 닫을 수 없습니다.: %v", err)
	}
	// -------- client가 자동 종료되도록 추가한 코드 --------

    // 모든 메시지를 보낸 후에 서버의 응답을 받을 때까지 기다립니다.
    for {
        response, err := stream.Recv()
        if err != nil {
            // 서버가 스트림을 닫으면 종료합니다.
            if err == io.EOF { //client의 스트림 종료 신호를 처리
                break
            }
            log.Fatalf("메시지를 수신할 수 없습니다.: %v", err)
        }
        fmt.Printf("[server to client] %s\n", response.GetMessage())
    }
}


func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("서버에 연결할 수 없습니다.: %v", err)
	}
	defer conn.Close()

	client := bidirectional.NewBidirectionalClient(conn)

	sendMessage(client)
}
