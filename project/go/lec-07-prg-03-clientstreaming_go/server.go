package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "project3/clientstreaming"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedClientStreamingServer
}

func (s *server) GetServerResponse(stream pb.ClientStreaming_GetServerResponseServer) error {
	fmt.Println("Server processing gRPC client-streaming.")
	count := 0

	for {
		_, err := stream.Recv()
		if err != nil {
			if err == context.Canceled {
				break
			}
			if err == context.DeadlineExceeded {
				return err
			}
			break
		}
		count++
	}

	return stream.SendAndClose(&pb.Number{Value: int32(count)})
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterClientStreamingServer(s, &server{})

	fmt.Println("Starting server. Listening on port 50051.")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
