package utils

import (
	"context"
	lrpb "labyrinth/pkg/proto"
	"log"
	"time"
)

type Client struct {
	Client lrpb.LabyrinthServiceClient
	X      int
	Y      int
	Health int
	Score  int
	Spells int
	Status string
}

func (c *Client) GetLabyrinthInfo() (int, int) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	info, err := c.Client.GetLabyrinthInfo(ctx, &lrpb.Empty{})
	if err != nil {
		log.Fatalf("error in getting servers %v\n", err)
	}
	return int(info.Height), int(info.Width)
}

func (c *Client) GetPlayerStatus() (int, int, int, int, int) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	status, err := c.Client.GetPlayerStatus(ctx, &lrpb.Empty{})
	if err != nil {
		log.Fatalf("error in getting status %v\n", err)
	}
	c.Score = int(status.Score)
	c.Health = int(status.Health)
	c.X = int(status.Position.X)
	c.Y = int(status.Position.Y)
	c.Spells = int(status.Spells)
	return c.Score, c.Health, c.X, c.Y, c.Spells
}

func (c *Client) RegisterMove(move rune) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var direction lrpb.Move_Direction

	switch move {
	case 'w':
		direction = lrpb.Move_UP
	case 's':
		direction = lrpb.Move_DOWN
	case 'a':
		direction = lrpb.Move_LEFT
	case 'd':
		direction = lrpb.Move_RIGHT
	default:
		log.Printf("Invalid move: %c", move)
		return
	}

	moveRequest := &lrpb.Move{
		Direction: direction,
	}

	resp, err := c.Client.RegisterMove(ctx, moveRequest)
	if err != nil {
		log.Printf("Error registering move: %v", err)
		return
	}
	switch resp.Status {
	case lrpb.MoveResponse_SUCCESS:
		c.Status = "SUCCESS"
	case lrpb.MoveResponse_FAILURE:
		c.Status = "FAILURE"
	case lrpb.MoveResponse_VICTORY:
		c.Status = "VICTORY"
	case lrpb.MoveResponse_DEATH:
		c.Status = "DEATH"
	default:
		log.Fatalf("Invalid move: %c", move)
		return
	}
}
