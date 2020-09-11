package chat

import (
	"context"
	"log"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("Recebida a mensagem do cliente: %s", message.Body)
	return &Message{Body: "Olá Servidor!"}, nil
}
