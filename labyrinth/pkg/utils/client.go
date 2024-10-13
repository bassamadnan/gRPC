package utils

import (
	"context"
	"errors"
	"fmt"
	"io"
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

func (c *Client) Revelio(X int, Y int, cellType string, grid [][]string) {
	if c.Spells == 0 {
		println("spells over")
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()
	var reqType lrpb.RevelioRequest_Tile
	switch cellType {
	case "E":
		reqType = lrpb.RevelioRequest_EMPTY
	case "W":
		reqType = lrpb.RevelioRequest_WALL
	case "C":
		reqType = lrpb.RevelioRequest_COIN
	}
	stream, err := c.Client.Revelio(ctx, &lrpb.RevelioRequest{
		Position: &lrpb.Position{
			X: int32(X),
			Y: int32(Y),
		},
		Type: reqType,
	})
	if err != nil {
		log.Fatalf("revelio error %v\n", err)
	}
	for {
		response, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.revelio(_) = _, %v", c, err)
		}
		x, y := int(response.Position.X), int(response.Position.Y)
		// log.Printf("recieved for %v, %v\n", x, y)
		grid[y][x] = cellType
	}
}

func (c *Client) Bombarda(grid [][]string) {
	stream, err := c.Client.Bombarda(context.Background())
	if err != nil {
		log.Fatalf("%v.Bombarda(_) = _, %v", c, err)
	}
	coords := make([][2]int, 3)

	for i := 0; i < 3; i++ {
		var x, y int
		fmt.Printf("Enter x y for point %d: ", i+1)
		_, err := fmt.Scanf("%d %d", &x, &y)
		if err != nil || grid[y][x] == "P" || grid[y][x] == "G" {
			log.Printf("err reading input: %v", err)
			i-- // retry
			continue
		}
		coords[i] = [2]int{x, y}
	}

	for _, coord := range coords {
		err := stream.Send(&lrpb.BombardaRequest{
			Position: &lrpb.Position{
				X: int32(coord[0]),
				Y: int32(coord[1]),
			},
		})
		if err != nil {
			log.Fatalf("Error sending coordinate: %v", err)
		}
		grid[coord[1]][coord[0]] = " " // incase we have seen it before as W
	}

	reply, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("%v.CloseAndRecv() got error %v, want %v", stream, err, nil)

	}

	if reply.Success {
		fmt.Println("Bombarda spell cast successfully!")
	} else {
		fmt.Println("Bombarda spell failed.")
	}
}
