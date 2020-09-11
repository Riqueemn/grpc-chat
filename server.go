package main

//go:generate protoc --go_out=plugins=grpc:chat chat.proto

import (
	"fmt"
	"log"
	"net"

	"github.com/riqueemn/grpc/chat"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Rodando gRPC")

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Falha ao escutar: %v", err)
	}

	s := chat.Server{}

	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Falha ao escutar: %v", err)
	}
}
