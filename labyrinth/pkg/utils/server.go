package utils

import (
	"context"
	"fmt"
	"io"
	lrpb "labyrinth/pkg/proto"

	"google.golang.org/grpc"
)

type Server struct {
	lrpb.UnimplementedLabyrinthServiceServer
	Grid   [][]string
	M      int // rows
	N      int // columns
	X      int // player position-x
	Y      int // player position-y
	Health int
	Spells int
	Score  int
}

// GetLabyrinthInfo(context.Context, *Empty) (*LabyrinthInfo, error)
// GetPlayerStatus(context.Context, *Empty) (*PlayerStatus, error)
// RegisterMove(context.Context, *Move) (*MoveResponse, error)
// Revelio(*RevelioRequest, grpc.ServerStreamingServer[RevelioResponse]) error
// Bombarda(grpc.ClientStreamingServer[BombardaRequest, BombardaResponse]) error
func (s *Server) GetLabyrinthInfo(ctx context.Context, req *lrpb.Empty) (*lrpb.LabyrinthInfo, error) {
	return &lrpb.LabyrinthInfo{
		Height: int32(s.M),
		Width:  int32(s.N),
	}, nil
}

func (s *Server) GetPlayerStatus(ctx context.Context, req *lrpb.Empty) (*lrpb.PlayerStatus, error) {
	return &lrpb.PlayerStatus{
		Position: &lrpb.Position{
			X: int32(s.X),
			Y: int32(s.Y),
		},
		Health: int32(s.Health),
		Spells: int32(s.Spells),
		Score:  int32(s.Score),
	}, nil
}

func (s *Server) RegisterMove(ctx context.Context, req *lrpb.Move) (*lrpb.MoveResponse, error) {
	direction := req.Direction
	fmt.Printf("Player moved to border")
	if !CheckBounds(direction, s.X, s.Y, s.M, s.N) {
		return &lrpb.MoveResponse{Status: lrpb.MoveResponse_FAILURE}, nil
	}

	if CheckTile(direction, s.X, s.Y, s.Grid, "W") {
		fmt.Printf("Player moved to wall")
		s.Health--
		status := lrpb.MoveResponse_FAILURE
		if s.Health == 0 {
			status = lrpb.MoveResponse_DEATH
		}
		return &lrpb.MoveResponse{Status: status, Wall: true}, nil
	}

	newX, newY, coinCollected := UpdatePlayerPosition(direction, s.X, s.Y, s.Grid)
	s.X, s.Y = newX, newY
	fmt.Printf("Player sucessfuly moved to :%v , %v", newX, newY)
	if coinCollected {
		s.Score++
		return &lrpb.MoveResponse{Status: lrpb.MoveResponse_SUCCESS}, nil
	}

	if s.Grid[newY][newX] == "G" {
		return &lrpb.MoveResponse{Status: lrpb.MoveResponse_VICTORY}, nil
	}

	return &lrpb.MoveResponse{Status: lrpb.MoveResponse_SUCCESS}, nil
}

func (s *Server) Revelio(req *lrpb.RevelioRequest, stream grpc.ServerStreamingServer[lrpb.RevelioResponse]) error {
	X, Y := int(req.Position.X), int(req.Position.Y)
	var cellType string
	switch req.Type {
	case lrpb.RevelioRequest_EMPTY:
		cellType = "E"
	case lrpb.RevelioRequest_COIN:
		cellType = "C"

	case lrpb.RevelioRequest_WALL:
		cellType = "W"
	}
	fmt.Printf("revelio req for  %v, %v, %v", X, Y, cellType)
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			newX, newY := X+i, Y+j
			if newX >= 0 && newX < s.N && newY >= 0 && newY < s.M {
				if s.Grid[newY][newX] == cellType {
					// fmt.Printf("sending %v %v\n", newX, newY)
					response := &lrpb.RevelioResponse{
						Position: &lrpb.Position{
							X: int32(newX),
							Y: int32(newY),
						},
					}
					if err := stream.Send(response); err != nil {
						return err
					}
				}
			}
		}
	}
	s.Spells--
	return nil
}

func (s *Server) Bombarda(stream grpc.ClientStreamingServer[lrpb.BombardaRequest, lrpb.BombardaResponse]) error {
	for {
		point, err := stream.Recv()
		if err == io.EOF {
			s.Spells--
			PrintGridAsTable(s.Grid) // show updated table
			return stream.SendAndClose(&lrpb.BombardaResponse{Success: true})
		}
		if err != nil {
			return err
		}
		x, y := int(point.Position.X), int(point.Position.Y)

		if x >= 0 && x < s.N && y >= 0 && y < s.M {
			s.Grid[y][x] = "E"
		} else {
			return stream.SendAndClose(&lrpb.BombardaResponse{Success: false})
		}
	}
}
