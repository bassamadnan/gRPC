package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	rspb "myuber/pkg/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type requestInfo struct {
	pickup      string
	destination string
	id          string
}

func RequestRide(ri *requestInfo, c rspb.RideServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req := &rspb.RideRequest{
		Pickup:      ri.pickup,
		Destination: ri.destination,
		RiderId:     ri.id,
	}
	rideResponse, err := c.RequestRide(ctx, req)
	if err != nil {
		log.Fatalf("error in getting servers %v\n", err)
	}
	fmt.Printf("got response: %v, status: %v, type %T\n", rideResponse, rideResponse.Success, rideResponse.Success)
}

func GetRideStatus(riderId string, c rspb.RideServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req := &rspb.RideStatusRequest{
		RiderId: riderId,
	}
	rideStatusResponse, err := c.GetRideStatus(ctx, req)
	if err != nil {
		log.Fatalf("error in getting servers %v\n", err)
	}
	if rideStatusResponse != nil {
		log.Fatalf("empty response in getridestatus\n")
	}
	fmt.Printf("got response: %v, status: %v, type %T\n", rideStatusResponse, rideStatusResponse.Status, rideStatusResponse.Status)
}

func main() {

	BASE_SERVER_ADDR := "localhost:5050"
	clientId := flag.Int("id", 1, "rider id")
	pickup := fmt.Sprintf("pickup%v", *clientId)
	destination := fmt.Sprintf("destination%v", *clientId)
	riderid := fmt.Sprintf("riderid%v", *clientId)
	info := &requestInfo{
		pickup:      pickup,
		destination: destination,
		id:          riderid,
	}

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	conn, err := grpc.NewClient(BASE_SERVER_ADDR, opts...)
	if err != nil {
		log.Fatalf("conn failed %v", err)
	}
	defer conn.Close()

	client := rspb.NewRideServiceClient(conn)
	print(client)
	RequestRide(info, client)

	for {

		var input string
		fmt.Print("Enter 1 to get ride/status: ")
		n, err := fmt.Scanf("%v", &input)
		if err != nil || n != 1 {
			log.Fatalf("scanf error %v", err)
		}

		if input == "1" {
			GetRideStatus(riderid, client)
		}
	}
}
