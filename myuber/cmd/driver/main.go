package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"sync"
	"time"

	"myuber/internal/auth"
	rscd "myuber/internal/client/driver" // ride sharing client; driver type
	utils "myuber/internal/utils"
	lbpb "myuber/pkg/proto/loadbalance"
	rspb "myuber/pkg/proto/rideshare"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const LOAD_BALANCER_ADDR string = "localhost:7070"

func getServersFromLoadBalancer(lbClient lbpb.LoadBalanceServiceClient) ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	resp, err := lbClient.GetServers(ctx, &lbpb.Empty{})
	if err != nil {
		log.Fatalf("failed to get servers from load balancer: %v", err)
	}
	return resp.Servers, nil
}

func getServerFromLoadBalancer(lbClient lbpb.LoadBalanceServiceClient) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	resp, err := lbClient.GetServer(ctx, &lbpb.Empty{})
	if err != nil {
		log.Fatalf("failed to get server from load balancer: %v", err)
	}
	return resp.Server, nil
}

func connectToServer(serverAddr string, driverID string, requestChan chan *rspb.DriverRideRequest) error {
	tlsCredentials, err := auth.DriverLoadTLSCredentials()
	if err != nil {
		log.Fatalf("cannot load TLS credentials: %v", err)
	}
	opts := []grpc.DialOption{grpc.WithTransportCredentials(tlsCredentials)}
	conn, err := grpc.NewClient(serverAddr, opts...)
	if err != nil {
		log.Fatalf("failed to connect to server %s: %v", serverAddr, err)
	}
	defer conn.Close()

	client := rspb.NewRideServiceClient(conn)
	stream, err := client.ConnectDriver(context.Background(), &rspb.DriverRequest{DriverId: driverID})
	if err != nil {
		log.Fatalf("stream error: %v", err)
	}

	for {
		request, err := stream.Recv()
		if err != nil {
			log.Fatalf("error receiving from stream: %v", err)
		}
		requestChan <- request
	}
}

func main() {
	clientId := flag.Int("id", 1, "driver id")
	flag.Parse()
	driverID := fmt.Sprintf("driver%v", *clientId)
	lbConn, err := grpc.NewClient(LOAD_BALANCER_ADDR, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to load balancer: %v", err)
	}
	defer lbConn.Close()
	lbClient := lbpb.NewLoadBalanceServiceClient(lbConn)
	servers, err := getServersFromLoadBalancer(lbClient)
	if err != nil {
		log.Fatalf("failed to get servers: %v", err)
	}
	requestChan := make(chan *rspb.DriverRideRequest, 100) // incoming ride requests from servers are stored here
	var wg sync.WaitGroup
	for _, serverAddr := range servers {
		wg.Add(1)
		go func(addr string) {
			defer wg.Done()
			for {
				err := connectToServer(addr, driverID, requestChan)
				if err != nil {
					log.Printf("Error connecting to server %s: %v. Retrying in 5 seconds...", addr, err)
					time.Sleep(5 * time.Second)
				}
			}
		}(serverAddr)
	}

	for {
		select {
		case request := <-requestChan:
			fmt.Printf("Received request: %v\n", request)
			input := utils.TakeInput("accept/reject")

			serverAddr, err := getServerFromLoadBalancer(lbClient)
			if err != nil {
				log.Printf("Failed to get server from load balancer: %v. Skipping this request.", err)
				continue
			}

			tlsCredentials, err := auth.DriverLoadTLSCredentials()
			if err != nil {
				log.Printf("DriverLoadTLSCredentials error %v\n.", err)
				continue
			}
			opts := []grpc.DialOption{grpc.WithTransportCredentials(tlsCredentials)}
			conn, err := grpc.NewClient(serverAddr, opts...)
			if err != nil {
				log.Printf("Failed to connect to server %s: %v. Skipping this request.", serverAddr, err)
				continue
			}
			client := rspb.NewRideServiceClient(conn)

			fmt.Printf("sending response to %v\n", serverAddr)
			if input == "1" {
				resp := rscd.AcceptRide(request.RiderId, driverID, client)
				if !resp.Success {
					fmt.Printf("rider: %v not waiting anymore\n", request.RiderId)
					conn.Close()
					continue
				}
				input := utils.TakeInput("complete")
				if input == "1" {
					rscd.CompleteRide(driverID, client)
				}
			} else if input == "0" {
				rscd.RejectRide(request.RiderId, driverID, client)
			}
			conn.Close()
		}
	}

}
