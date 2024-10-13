package main

import (
	"context"
	dpb "docsbackend/proto"
	"fmt"
	"io"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

type Server struct {
	dpb.UnimplementedDocumentServiceServer
	clients map[string]bool
}

func (s *Server) SendRecvNumbers(stream grpc.BidiStreamingServer[dpb.Number, dpb.Number]) error {
	fmt.Print("called\n")
	recvNumbers := make([]int, 0, 5)
	sentNumbers := []int{7, 3, 1, 3, 1}

	for {
		in, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		recvNumbers = append(recvNumbers, int(in.Number))
	}

	for _, num := range sentNumbers {
		time.Sleep(1 * time.Second)
		fmt.Printf("sent %v\n", num)
		if err := stream.Send(&dpb.Number{Number: int32(num)}); err != nil {
			return err
		}
	}

	fmt.Print(recvNumbers)
	return nil
}
func (s *Server) SayHello(ctx context.Context, req *dpb.HelloRequest) (*dpb.HelloReply, error) {
	name := req.Name
	fmt.Printf("Received request from %v\n", name)
	response := fmt.Sprintf("server: hello %s", name)
	return &dpb.HelloReply{Message: response}, nil
}

func main() {
	// Listen on port 5050
	lis, err := net.Listen("tcp", ":5050")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	dpb.RegisterDocumentServiceServer(grpcServer, &Server{
		clients: make(map[string]bool),
	})

	log.Printf("Starting standard gRPC server on :5050")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
