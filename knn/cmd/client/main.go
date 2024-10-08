package main

import (
	"context"
	"fmt"
	knnpb "knn/pkg/api"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func getServers(client knnpb.KNNClient) int {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	servers, err := client.GetServers(ctx, &knnpb.Empty{})
	if err != nil {
		log.Fatalf("error in getting servers %v\n", err)
	}
	return int(servers.Servers)
}

func main() {
	BASE_SERVER_ADDR := "localhost:5051" // first server always at 5051
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.NewClient(BASE_SERVER_ADDR, opts...)
	if err != nil {
		log.Fatalf("conn failed %v", err)
	}

	client := knnpb.NewKNNClient(conn)

	num_servers := getServers(client)
	fmt.Printf("Number of servers %v \n", num_servers)
	conn.Close()
}
