package main

import (
	"context"
	"docs/cmd/client/editor"
	mpb "docs/pkg/proto/message"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	Client mpb.MessageServiceClient
	Name   string
}

func (c *Client) sayHello() string {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	info, err := c.Client.SendMessage(ctx, &mpb.Message{Body: c.Name})
	if err != nil {
		log.Fatalf("error in getting servers %v\n", err)
	}
	return info.Body
}

func main() {
	BASE_SERVER_ADDR := "localhost:5050"
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	conn, err := grpc.NewClient(BASE_SERVER_ADDR, opts...)
	if err != nil {
		log.Fatalf("conn failed %v", err)
	}
	defer conn.Close()
	client := Client{
		Client: mpb.NewMessageServiceClient(conn),
		Name:   "client1",
	}
	uiConfig := UIConfig{
		EditorConfig: editor.EditorConfig{
			ScrollEnabled: true,
		},
	}
	err2 := initUI(&client, uiConfig)
	if err2 != nil {
		log.Fatalf("init ui error %v\n", err)
	}
}
