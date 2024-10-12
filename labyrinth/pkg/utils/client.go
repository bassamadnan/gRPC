package utils

import (
	"context"
	"errors"
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

func (c *Client) RegisterMove(move rune) (bool, error) {
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
		return false, errors.New("invalid move error")
	}

	moveRequest := &lrpb.Move{
		Direction: direction,
	}

	resp, err := c.Client.RegisterMove(ctx, moveRequest)
	if err != nil {
		log.Printf("Error registering move: %v", err)
		return false, err
	}

	var wall bool

	switch resp.Status {
	case lrpb.MoveResponse_SUCCESS:
		c.Status = "SUCCESS"
	case lrpb.MoveResponse_FAILURE:
		c.Status = "FAILURE"
		wall = resp.Wall
	case lrpb.MoveResponse_VICTORY:
		c.Status = "VICTORY"
	case lrpb.MoveResponse_DEATH:
		c.Status = "DEATH"
	default:
		log.Fatalf("Invalid move: %c", move)
		return false, errors.New("invalid move error")
	}
	return wall, nil
}

func (c *Client) HandleMove(char rune, grid [][]string) {
	currX, currY := c.X, c.Y
	wall, err := c.RegisterMove(char)
	if err != nil {
		log.Printf("Error at register move: %v\n", err)
		return
	}

	if c.Status == "SUCCESS" || c.Status == "VICTORY" {
		grid[currY][currX] = " "
	}

	if wall {
		wallX, wallY := c.X, c.Y
		switch char {
		case 'w':
			wallY--
		case 'a':
			wallX--
		case 's':
			wallY++
		case 'd':
			wallX++
		}
		if wallX >= 0 && wallX < len(grid[0]) && wallY >= 0 && wallY < len(grid) {
			grid[wallY][wallX] = "W"
		}
	}
}
