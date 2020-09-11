package main

import (
	"context"
	"log"

	"github.com/riqueemn/grpc/chat"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Não conectou: %s", err)
	}
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)

	message := chat.Message{
		Body: "Olá cliente",
	}

	response, err := c.SayHello(context.Background(), &message)
	if err != nil {
		log.Fatalf("Erro no SayHello: %s", err)
	}

	log.Printf("Respondendo para o Servidor: %s", response.Body)
}
