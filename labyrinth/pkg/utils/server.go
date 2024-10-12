package utils

import (
	"context"
	"fmt"
	lrpb "labyrinth/pkg/proto"
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
