package main

import (
	"fmt"
	lrpb "labyrinth/pkg/proto"
	utils "labyrinth/pkg/utils"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	grid, err := utils.LoadGridFromFile("grid.txt")
	if err != nil {
		fmt.Println("Error loading grid:", err)
		return
	}

	utils.PrintGridAsTable(grid)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 5050))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	lrpb.RegisterLabyrinthServiceServer(s, &utils.Server{
		Grid:   grid,
		M:      len(grid),
		N:      len(grid[0]),
		X:      0,
		Y:      0,
		Health: 3,
		Spells: 3,
		Score:  0,
	})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
