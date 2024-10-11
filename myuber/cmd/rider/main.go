package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	auth "myuber/internal/auth"
	rscr "myuber/internal/client/rider" //rider sharing client: rider type
	lbpb "myuber/pkg/proto/loadbalance"
	rspb "myuber/pkg/proto/rideshare"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const LOAD_BALANCER_ADDR string = "localhost:7070"

func getServerFromLoadBalancer(lbClient lbpb.LoadBalanceServiceClient) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	server, err := lbClient.GetServer(ctx, &lbpb.Empty{})
	if err != nil {
		return "", fmt.Errorf("failed to get server from load balancer: %v", err)
	}
	return server.Server, nil
}

func main() {
	clientId := flag.Int("id", 1, "rider id")
	flag.Parse()
	pickup := fmt.Sprintf("pickup%v", *clientId)
	destination := fmt.Sprintf("destination%v", *clientId)
	riderid := fmt.Sprintf("riderid%v", *clientId)

	info := &rscr.SelfInfo{
		Pickup:      pickup,
		Destination: destination,
		Id:          riderid,
	}

	lbConn, err := grpc.NewClient(LOAD_BALANCER_ADDR, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to load balancer: %v", err)
	}
	defer lbConn.Close()
	lbClient := lbpb.NewLoadBalanceServiceClient(lbConn)

	tlsCredentials, err := auth.RiderLoadTLSCredentials()
	if err != nil {
		log.Fatal("cannot load TLS credentials: ", err)
	}

	// Request ride once, outside the loop
	serverAddr, err := getServerFromLoadBalancer(lbClient)
	if err != nil {
		log.Fatalf("load balance error: %v", err)
	}
	opts := []grpc.DialOption{grpc.WithTransportCredentials(tlsCredentials)}
	conn, err := grpc.NewClient(serverAddr, opts...)
	if err != nil {
		log.Fatalf("conn failed %v", err)
	}
	client := rspb.NewRideServiceClient(conn)
	rscr.RequestRide(info, client)
	conn.Close()
	fmt.Printf("client %v request to %v\n", clientId, serverAddr)
	for {
		var input string
		fmt.Print("Enter 1 to get ride/status: ")
		n, err := fmt.Scanf("%v", &input)
		if err != nil || n != 1 {
			log.Fatalf("scanf error %v", err)
		}

		if input == "1" {
			serverAddr, err := getServerFromLoadBalancer(lbClient)
			if err != nil {
				log.Printf("load balance error: %v. Retrying in 5 seconds...", err)
				time.Sleep(5 * time.Second)
				continue
			}
			conn, err := grpc.NewClient(serverAddr, opts...)
			if err != nil {
				log.Printf("conn failed %v. Retrying...", err)
				continue
			}
			client := rspb.NewRideServiceClient(conn)
			rscr.GetRideStatus(riderid, client)
			conn.Close()
			fmt.Printf("client %v request to %v\n", clientId, serverAddr)
		}
	}
}
