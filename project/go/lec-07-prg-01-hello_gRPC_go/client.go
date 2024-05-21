package main

import (
	"context"
	"fmt"
	"log"

	pb "project1/hello_grpc"

	"google.golang.org/grpc"
)

func main() {
	// gRPC 서버에 연결하는 코드
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()

	// 서버에서 제공하는 클라이언트 인스턴스 생성
	client := pb.NewMyServiceClient(conn)

	// 원격 함수 호출에 필요한 요청 메시지 생성
	request := &pb.MyNumber{Value: 4}

	// 원격 함수 호출 및 결과 처리
	response, err := client.MyFunction(context.Background(), request)
	if err != nil {
		log.Fatalf("Error calling MyFunction: %v", err)
	}

	// 결과 출력
	fmt.Println("gRPC result:", response.Value)
}
