package main

import (
	"context"
	dpb "docsbackend/proto"
	"fmt"
	"log"
	"net/http"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc"
)

type Server struct {
	dpb.UnimplementedGreeterServer
}

func (s *Server) SayHello(ctx context.Context, req *dpb.HelloRequest) (*dpb.HelloReply, error) {
	name := req.Name
	fmt.Printf("recieved request from %v\n", name)
	response := fmt.Sprintf("server: hello %s", name)
	return &dpb.HelloReply{Message: response}, nil
}

func main() {
	grpcServer := grpc.NewServer()
	dpb.RegisterGreeterServer(grpcServer, &Server{})

	wrappedGrpc := grpcweb.WrapServer(grpcServer,
		grpcweb.WithOriginFunc(func(origin string) bool {
			return true // cors
		}))

	handler := http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		resp.Header().Set("Access-Control-Allow-Origin", "*")
		resp.Header().Set("Access-Control-Allow-Headers", "Content-Type,x-grpc-web,x-user-agent")
		if req.Method == "OPTIONS" {
			return
		}
		wrappedGrpc.ServeHTTP(resp, req)
	})

	httpServer := &http.Server{
		Addr:    ":5050",
		Handler: handler,
	}

	log.Printf("hosting gRPC-Web server on :5050")
	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
