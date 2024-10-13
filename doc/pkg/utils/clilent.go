package utils

import (
	"context"
	mpb "docs/pkg/proto/message"
	"log"
	"time"
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
