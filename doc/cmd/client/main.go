package main

import (
	"context"
	"docs/cmd/client/editor"
	"docs/crdt"
	dpb "docs/pkg/proto/docs"
	"docs/pkg/utils"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	Client dpb.DocsServiceClient
	Name   string
	Stream dpb.DocsService_EditDocClient
}

var (
	doc     = crdt.New()
	e       = editor.NewEditor(editor.EditorConfig{})
	logger  = logrus.New()
	client  = Client{}
	msgChan = make(chan *dpb.Message)
)

func (c *Client) sendMessage(message *dpb.Message) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := c.Client.SendMessage(ctx, message)
	if err != nil {
		log.Fatalf("error in getting servers %v\n", err)
		return nil
	}
	return nil
}

func (c *Client) registerClient() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	message := &dpb.Message{Username: c.Name}
	docRecieved, err := c.Client.RegisterClient(ctx, message)
	if err != nil {
		// log.Fatalf("error in getting servers %v\n", err)
		return nil
	}
	doc = *utils.GetDocument(docRecieved)
	e.SetText(crdt.Content(doc))
	return nil
}
func (c *Client) sendError() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := c.Client.SendError(ctx, &dpb.Message{Username: c.Name})
	if err != nil {
		log.Fatalf("error in getting servers %v\n", err)
		return nil
	}
	return nil
}

func (c *Client) startCollab() error {
	stream, err := c.Client.EditDoc(context.Background())
	if err != nil {
		log.Fatalf("stream error %v", err)
		return err
	}
	c.Stream = stream
	return nil
}

func (c *Client) receiveMessages(msgChan chan<- *dpb.Message) {
	for {
		msg, err := c.Stream.Recv()
		if err != nil {
			log.Fatalf("stream recv error %v", err)
			close(msgChan)
			return
		}
		msgChan <- msg
	}
}
func main() {
	id := flag.Int("id", 1, "client id")
	flag.Parse()
	BASE_SERVER_ADDR := "localhost:5050"
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	conn, err := grpc.NewClient(BASE_SERVER_ADDR, opts...)
	if err != nil {
		log.Fatalf("conn failed %v", err)
	}
	defer conn.Close()
	// crdt.IsCRDT(&doc)
	client = Client{
		Client: dpb.NewDocsServiceClient(conn),
		Name:   fmt.Sprintf("client%v", *id),
	}

	uiConfig := UIConfig{
		EditorConfig: editor.EditorConfig{
			ScrollEnabled: true,
		},
	}
	err = client.startCollab()
	if err != nil {
		client.sendError()
		return
	}
	go client.receiveMessages(msgChan)
	client.registerClient()
	err2 := initUI(uiConfig)
	if err2 != nil {
		log.Fatalf("init ui error %v\n", err)
	}

}
