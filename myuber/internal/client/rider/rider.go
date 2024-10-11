package client

import (
	"context"
	"fmt"
	"log"
	rspb "myuber/pkg/proto/rideshare"
	"time"
)

type SelfInfo struct {
	Pickup      string
	Destination string
	Id          string
}

func RequestRide(ri *SelfInfo, c rspb.RideServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req := &rspb.RideRequest{
		Pickup:      ri.Pickup,
		Destination: ri.Destination,
		RiderId:     ri.Id,
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
	if rideStatusResponse == nil {
		log.Fatalf("empty response in getridestatus\n")
	}
	fmt.Printf("got response: %v, status: %v, type %T\n", rideStatusResponse, rideStatusResponse.Status, rideStatusResponse.Status)
}
