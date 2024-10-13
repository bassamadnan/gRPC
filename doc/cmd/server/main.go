package main

import (
	"context"
	dpb "docs/pkg/proto/docs"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	dpb.UnsafeDocsServiceServer
}

// SendMessage(context.Context, *Message) (*MessageResponse, error)

func (s *Server) SendMessage(ctx context.Context, msg *dpb.Message) (*dpb.MessageResponse, error) {
	fmt.Printf("\nReceived msg: {\n"+
		"  Document: %v, "+
		"  Text: %v,\n"+
		"  Username: %v,"+
		"  MessageType: %v,"+
		"  Operation: %v"+
		"}\n",
		msg.Document, msg.Text, msg.Username, msg.MessageType, msg.Operation)
	return &dpb.MessageResponse{
		Success: true,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 5050))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	dpb.RegisterDocsServiceServer(s, &Server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
