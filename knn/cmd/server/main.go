package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math"
	"net"

	"google.golang.org/grpc"

	"knn/internal/knn"
	"knn/internal/utils"
	knnpb "knn/pkg/api"
)

type server struct {
	knnpb.UnimplementedKNNServer
	dataset      []knn.Point
	serverNumber int32
	totalServers int32
}

func (s *server) GetKNN(ctx context.Context, query *knnpb.Query) (*knnpb.Distances, error) {
	queryPoint := knn.Point{X: query.Point.X, Y: query.Point.Y}
	k := int(query.K)

	nearestNeighbors := utils.GetKNN(uint16(k), queryPoint, s.dataset)

	var distances knnpb.Distances
	for _, neighbor := range nearestNeighbors {
		distances.Points = append(distances.Points, &knnpb.Points{
			Point:    &knnpb.Point{X: neighbor.Point.X, Y: neighbor.Point.Y},
			Distance: neighbor.Distance,
		})
	}

	return &distances, nil
}

func (s *server) GetServers(ctx context.Context, empty *knnpb.Empty) (*knnpb.Servers, error) {
	return &knnpb.Servers{Servers: s.totalServers}, nil
}

func main() {
	BASE_PORT := 5050
	// flags
	serverNumber := flag.Int("server", 0, "The server number (starting from 0)")
	totalServers := flag.Int("total", 1, "Total number of servers")

	flag.Parse()

	if *serverNumber < 0 || *serverNumber >= *totalServers {
		log.Fatalf("Invalid server number. Must be between 0 and %d", *totalServers-1)
	}

	port := BASE_PORT + *serverNumber + 1

	// total number of lines in the file
	totalLines, err := utils.CountFileLines("data.txt")
	if err != nil {
		log.Fatalf("Failed to count lines in dataset: %v", err)
	}

	// chunk , start and end lines
	linesPerServer := int(math.Ceil(float64(totalLines) / float64(*totalServers)))
	start := *serverNumber*linesPerServer + 1
	end := (*serverNumber + 1) * linesPerServer
	if end > totalLines {
		end = totalLines
	}

	dataset, err := utils.ReadFile("data.txt", start, end)
	if err != nil {
		log.Fatalf("Failed to read dataset: %v", err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	knnpb.RegisterKNNServer(s, &server{
		dataset:      dataset,
		serverNumber: int32(*serverNumber),
		totalServers: int32(*totalServers),
	})

	fmt.Printf("Server %d of %d is running on :%d\n", *serverNumber+1, *totalServers, port)
	fmt.Printf("Handling lines %d to %d of %d total lines\n", start, end, totalLines)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
