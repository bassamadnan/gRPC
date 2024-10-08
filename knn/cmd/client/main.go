package main

import (
	"context"
	"flag"
	"fmt"
	knnpb "knn/pkg/api"
	"log"
	"sort"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const BASE_PORT = 5050

func getServers(client knnpb.KNNClient) int {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	servers, err := client.GetServers(ctx, &knnpb.Empty{})
	if err != nil {
		log.Fatalf("error in getting servers %v\n", err)
	}
	return int(servers.Servers)
}

func queryServer(serverAddr string, query *knnpb.Query, results chan<- []*knnpb.Points, wg *sync.WaitGroup) {
	defer wg.Done()

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	conn, err := grpc.NewClient(serverAddr, opts...)
	if err != nil {
		log.Printf("Failed to connect to server %s: %v", serverAddr, err)
		return
	}
	defer conn.Close()

	client := knnpb.NewKNNClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := client.GetKNN(ctx, query)
	if err != nil {
		log.Printf("Failed to get KNN from server %s: %v", serverAddr, err)
		return
	}

	results <- resp.Points
}

func main() {
	x := flag.Float64("x", 0, "X coordinate of the query point")
	y := flag.Float64("y", 0, "Y coordinate of the query point")
	k := flag.Int("k", 3, "Number of nearest neighbors to find")
	flag.Parse()

	BASE_SERVER_ADDR := fmt.Sprintf("localhost:%d", BASE_PORT+1) // first server always at 5051
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	conn, err := grpc.NewClient(BASE_SERVER_ADDR, opts...)
	if err != nil {
		log.Fatalf("conn failed %v", err)
	}
	defer conn.Close()

	client := knnpb.NewKNNClient(conn)

	num_servers := getServers(client)
	fmt.Printf("Number of servers %v \n", num_servers)

	query := &knnpb.Query{
		Point: &knnpb.Point{X: *x, Y: *y},
		K:     int32(*k),
	}

	results := make(chan []*knnpb.Points, num_servers)

	var wg sync.WaitGroup

	for i := 0; i < num_servers; i++ {
		wg.Add(1)
		serverAddr := fmt.Sprintf("localhost:%d", BASE_PORT+i+1)
		go queryServer(serverAddr, query, results, &wg)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	var allPoints []*knnpb.Points
	for serverPoints := range results {
		allPoints = append(allPoints, serverPoints...)
	}

	sort.Slice(allPoints, func(i, j int) bool {
		return allPoints[i].Distance < allPoints[j].Distance
	})

	fmt.Printf("The %d nearest neighbors to point (%.2f, %.2f) are:\n", *k, *x, *y)
	for i := 0; i < *k && i < len(allPoints); i++ {
		point := allPoints[i]
		fmt.Printf("%d. (%.2f, %.2f) - Distance: %.4f\n", i+1, point.Point.X, point.Point.Y, point.Distance)
	}
}
