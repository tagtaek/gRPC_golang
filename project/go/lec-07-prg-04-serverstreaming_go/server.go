package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "project2/serverstreaming" // 프로토 파일에서 생성된 Go 코드의 패키지 경로로 수정해야 합니다
)

type server struct {
	pb.UnimplementedServerStreamingServer
}

func (s *server) GetServerResponse(req *pb.Number, stream pb.ServerStreaming_GetServerResponseServer) error {
	message := []*pb.Message{
		{Message: "message #1"},
		{Message: "message #2"},
		{Message: "message #3"},
		{Message: "message #4"},
		{Message: "message #5"},
	}

	fmt.Printf("Server processing gRPC server-streaming {%d}.\n", req.Value)
	for _, msg := range message {
		if err := stream.Send(msg); err != nil {
			return status.Errorf(codes.Internal, "failed to send message: %v", err)
		}
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterServerStreamingServer(s, &server{})
	fmt.Println("Starting server. Listening on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
