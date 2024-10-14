package main

import (
	"context"
	crdt "docs/crdt"
	dpb "docs/pkg/proto/docs"
	utils "docs/pkg/utils"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"sync"

	"google.golang.org/grpc"
)

type Server struct {
	dpb.UnimplementedDocsServiceServer
	Clients  []string
	Active   map[string]bool
	Mu       sync.Mutex
	Document crdt.Document // does server store the Document? for now yes
	Streams  map[string]dpb.DocsService_EditDocServer
	Count    int
	LogFile  string
}

// SendMessage(context.Context, *Message) (*MessageResponse, error)
// RegisterClient(context.Context, *Message) (*Document, error)
// SendError(context.Context, *Message) (*Empty, error)
// EditDoc(grpc.BidiStreamingServer[Message, Message]) error

func (s *Server) EditDoc(stream dpb.DocsService_EditDocServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		// time.Sleep(500 * time.Millisecond)
		client := in.Username

		s.Mu.Lock()
		if _, exists := s.Streams[client]; !exists {
			s.Streams[client] = stream
			s.Active[client] = true
		}
		if in.MessageType == dpb.Message_JOIN {
			fmt.Printf("stream created %v\n", in.Username)
			// dont do anything, just create above stream
		} else {
			s.processMessage(in)
			s.forwardMessageToClients(in, client)
		}
		s.Mu.Unlock()
	}
}

func (s *Server) processMessage(msg *dpb.Message) {
	// s.Mu.Lock()
	// defer s.Mu.Unlock()

	fmt.Printf("\nProcessing msg: {\n"+
		"  Document: %v,\n"+
		"  Text: %v,\n"+
		"  Username: %v,\n"+
		"  MessageType: %v,\n"+
		"  Operation: %v\n"+
		"}\n",
		msg.Document, msg.Text, msg.Username, msg.MessageType, msg.Operation)

	// append the msg operation operation type, msg username , the index if insert then the text
	logEntry := fmt.Sprintf("Operation: %v, Username: %v", msg.Operation.OperationType, msg.Username)
	if msg.Operation.OperationType == dpb.Operation_INSERT {
		logEntry += fmt.Sprintf(", Index: %d, Text: %s", msg.Operation.Position, msg.Operation.Value)
	} else if msg.Operation.OperationType == dpb.Operation_DELETE {
		logEntry += fmt.Sprintf(", Index: %d", msg.Operation.Position)
	}

	f, err := os.OpenFile(s.LogFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("Error opening log file: %v\n", err)
	} else {
		defer f.Close()
		if _, err := f.WriteString(logEntry + "\n"); err != nil {
			fmt.Printf("Error writing to log file: %v\n", err)
		}
	}

	switch msg.Operation.OperationType {
	case dpb.Operation_INSERT:
		// fmt.Print("inserted")
		s.Document.Insert(int(msg.Operation.Position), msg.Operation.Value)
	case dpb.Operation_DELETE:
		// fmt.Print("deleted")
		s.Document.Delete(int(msg.Operation.Position))
	}
	fmt.Printf("\n\nCurrent docuemtn %v\n", s.Document)
	// s.Document = *utils.GetDocument(msg.Document) // update the server document
}

func (s *Server) forwardMessageToClients(msg *dpb.Message, sender string) {
	// s.Mu.Lock()
	// defer s.Mu.Unlock()
	fmt.Print(sender)
	msg.Document = utils.GetDocumentProto(s.Document)
	for client, stream := range s.Streams {
		if s.Active[client] {
			fmt.Print("sending to anoter client")
			err := stream.Send(msg)
			if err != nil {
				fmt.Printf("Error sending message to client %s: %v\n", client, err)
				s.Active[client] = false
			}
		}
	}
}

func (s *Server) RegisterClient(ctx context.Context, msg *dpb.Message) (*dpb.Document, error) {
	if s.Count == 0 {
		s.Count++
		log.Printf("first client sent")
		return nil, errors.New("first client")
	}
	s.Count++
	doc := utils.GetDocumentProto(s.Document)
	fmt.Printf("cleint %v sent -> %v", msg.Username, s.Document)
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

	s.Document = *utils.GetDocument(msg.Document) // this is our document now

	return &dpb.MessageResponse{
		Success: true,
	}, nil
}

func (s *Server) SendError(ctx context.Context, msg *dpb.Message) (*dpb.Empty, error) {
	fmt.Printf("Client %v error-ed\n", msg.Username)
	return &dpb.Empty{}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 5050))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	dpb.RegisterDocsServiceServer(s, &Server{
		Clients:  make([]string, 0),
		Active:   make(map[string]bool),
		Streams:  make(map[string]dpb.DocsService_EditDocServer),
		Count:    0,
		Document: crdt.New(),
		LogFile:  "logs.txt",
	})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
