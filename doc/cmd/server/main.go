package main

import (
	"context"
	crdt "docs/crdt"
	dpb "docs/pkg/proto/docs"
	utils "docs/pkg/utils"
	"errors"
	"fmt"
	"log"
	"net"
	"sync"

	"google.golang.org/grpc"
)

type Server struct {
	dpb.UnsafeDocsServiceServer
	Clients  []string
	Active   []bool
	Mu       sync.Mutex
	Document crdt.Document // does server store the Document? for now yes
}

// SendMessage(context.Context, *Message) (*MessageResponse, error)
// RegisterClient(context.Context, *Message) (*Document, error)

func (s *Server) SendError(ctx context.Context, msg *dpb.Message) (*dpb.Empty, error) {
	fmt.Print("Client %v error-ed\n", msg.Username)
	return &dpb.Empty{}, nil
}
func (s *Server) RegisterClient(ctx context.Context, msg *dpb.Message) (*dpb.Document, error) {
	if len(s.Document.Characters) == 0 {
		return nil, errors.New("first client")
	}
	doc := utils.GetDocumentProto(s.Document)
	return doc, nil
}

func (s *Server) SendMessage(ctx context.Context, msg *dpb.Message) (*dpb.MessageResponse, error) {
	fmt.Printf("\nReceived msg: {\n"+
		"  Document: %v,\n"+
		"  Text: %v,\n"+
		"  Username: %v,\n"+
		"  MessageType: %v,\n"+
		"  Operation: %v\n"+
		"}\n",
		msg.Document, msg.Text, msg.Username, msg.MessageType, msg.Operation)

	s.Document = *utils.GetDocument(msg.Document)

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
