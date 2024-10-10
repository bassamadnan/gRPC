package main

import (
	"context"
	"fmt"
	"log"
	"time"

	rspb "myuber/pkg/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func RequestRide(c rspb.RideServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	req := &rspb.RideRequest{
		Pickup:      "pickup1",
		Destination: "destination1",
		RiderId:     "riderid1",
	}
	rideStatusResponse, err := c.RequestRide(ctx, req)
	if err != nil {
		log.Fatalf("error in getting servers %v\n", err)
	}
	fmt.Printf("got response: %v, status: %v, type %T\n", rideStatusResponse, rideStatusResponse.Status, rideStatusResponse.Status)
}

func main() {
	BASE_SERVER_ADDR := "localhost:5050"
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	conn, err := grpc.NewClient(BASE_SERVER_ADDR, opts...)
	if err != nil {
		log.Fatalf("conn failed %v", err)
	}
	defer conn.Close()

	client := rspb.NewRideServiceClient(conn)
	print(client)
	RequestRide(client)
}
