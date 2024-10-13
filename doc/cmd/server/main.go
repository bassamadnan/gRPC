package main

import (
	"context"
	mpb "docs/pkg/proto/message"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	mpb.UnimplementedMessageServiceServer
}

// SendMessage(context.Context, *Message) (*Message, error)

func (s *Server) SendMessage(ctx context.Context, req *mpb.Message) (*mpb.Message, error) {
	body := req.Body
	fmt.Printf("recieved from %v\n", body)
	return &mpb.Message{
		Body: fmt.Sprintf("server: Hi %v!", body),
	}, nil

}
func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 5050))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	mpb.RegisterMessageServiceServer(s, &Server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
