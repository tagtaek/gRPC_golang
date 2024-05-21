package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "project1/hello_grpc"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedMyServiceServer
}

func (s *server) MyFunction(ctx context.Context, in *pb.MyNumber) (*pb.MyNumber, error) {
	return &pb.MyNumber{Value: in.Value * in.Value}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMyServiceServer(s, &server{})
	fmt.Println("Server is listening on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
