package main

import (
	"fmt"
	"io"
	"log"
	"net"

	pb "project4/bidirectional"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedBidirectionalServer
}

func (s *server) GetServerResponse(stream pb.Bidirectional_GetServerResponseServer) error {
	fmt.Println("Server processing gRPC bidirectional streaming.")
	for {
		message, err := stream.Recv()

		// -------- client가 자동 종료되도록 추가한 코드 --------
		if err == io.EOF {
			// 클라이언트가 스트림을 닫았음을 감지하고 스트림을 종료합니다.
			return nil
		}
		// -------- client가 자동 종료되도록 추가한 코드 --------
		//=> clientrk 스트림 닫았을때 처리를 올바르게 하도록 server.go 도 수정필요함.

		if err != nil {
			return err
		}
		if err := stream.Send(message); err != nil {
			return err
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterBidirectionalServer(s, &server{})
	fmt.Println("Starting server. Listening on port 50051.")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
