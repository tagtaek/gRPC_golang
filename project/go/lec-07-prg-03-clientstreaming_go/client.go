package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "project3/clientstreaming"

	"google.golang.org/grpc"
)

func generateMessages() <-chan *pb.Message {
	msgChan := make(chan *pb.Message)

	go func() {
		defer close(msgChan)

		messages := []*pb.Message{
			{Message: "message #1"},
			{Message: "message #2"},
			{Message: "message #3"},
			{Message: "message #4"},
			{Message: "message #5"},
		}

		for _, msg := range messages {
			fmt.Printf("[client to server] %s\n", msg.Message)
			msgChan <- msg
		}
	}()

	return msgChan
}

func sendMessage(client pb.ClientStreamingClient, ctx context.Context) {
	stream, err := client.GetServerResponse(ctx)
	if err != nil {
		log.Fatalf("Error creating stream: %v", err)
	}

	for msg := range generateMessages() {
		if err := stream.Send(msg); err != nil {
			log.Fatalf("Error sending message: %v", err)
		}
	}

	response, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error receiving response: %v", err)
	}

	fmt.Printf("[server to client] %d\n", response.Value)
}

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial server: %v", err)
	}
	defer conn.Close()

	client := pb.NewClientStreamingClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	sendMessage(client, ctx)
}
