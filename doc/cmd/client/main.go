package main

import (
	"docs/cmd/client/editor"
	mpb "docs/pkg/proto/message"
	utils "docs/pkg/utils"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	BASE_SERVER_ADDR := "localhost:5050"
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	conn, err := grpc.NewClient(BASE_SERVER_ADDR, opts...)
	if err != nil {
		log.Fatalf("conn failed %v", err)
	}
	defer conn.Close()
	client := utils.Client{
		Client: mpb.NewMessageServiceClient(conn),
		Name:   "client1",
	}
	uiConfig := utils.UIConfig{
		EditorConfig: editor.EditorConfig{
			ScrollEnabled: true,
		},
	}
	err2 := client.InitUI(uiConfig)
	if err2 != nil {
		log.Fatalf("init ui error %v\n", err)
	}
}
