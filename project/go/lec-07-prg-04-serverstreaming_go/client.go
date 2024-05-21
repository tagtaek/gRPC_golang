package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"

	pb "project2/serverstreaming" // 프로토 파일에서 생성된 Go 코드의 패키지 경로로 수정해야 합니다
)

func receiveMessage(client pb.ServerStreamingClient) {
	req := &pb.Number{Value: 5}
	stream, err := client.GetServerResponse(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling GetServerResponse: %v", err)
	}
	for {
		resp, err := stream.Recv()
		if err != nil {
			break
		}
		fmt.Printf("[server to client] %s\n", resp.Message)
	}
}

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewServerStreamingClient(conn)
	receiveMessage(client)
}
