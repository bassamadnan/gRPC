package client

import (
	"context"
	"fmt"
	"log"
	rspb "myuber/pkg/proto"
	"time"
)

func AcceptRide(riderId string, driverId string, c rspb.RideServiceClient) *rspb.AcceptRideResponse {
	req := &rspb.AcceptRideRequest{
		RiderId:  riderId,
		DriverId: driverId,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	acceptRideResponse, err := c.AcceptRide(ctx, req)
	if err != nil {
		log.Fatalf("error in getting servers %v\n", err)
	}
	if acceptRideResponse == nil {
		log.Fatalf("empty response in acceptRide\n")
	}
	fmt.Printf("got response: %v, status: %v, type %T\n", acceptRideResponse, acceptRideResponse.Success, acceptRideResponse.Success)
	return acceptRideResponse
}

func RejectRide(riderId string, driverId string, c rspb.RideServiceClient) {
	req := &rspb.RejectRideRequest{
		RiderId:  riderId,
		DriverId: driverId,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	rejectRideResponse, err := c.RejectRide(ctx, req)
	if err != nil {
		log.Fatalf("error in getting servers %v\n", err)
	}
	if rejectRideResponse == nil {
		log.Fatalf("empty response in rejectRide\n")
	}

	if rejectRideResponse.Success == false {
		log.Fatalf("could not reject ride ?? ")
	}
	fmt.Printf("got response: %v, status: %v, type %T\n", rejectRideResponse, rejectRideResponse.Success, rejectRideResponse.Success)
}

func CompleteRide(driverId string, c rspb.RideServiceClient) {
	req := &rspb.RideCompletionRequest{
		DriverId: driverId,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	rideCompletionResponse, err := c.CompleteRide(ctx, req)
	if err != nil {
		log.Fatalf("error in getting servers %v\n", err)
	}
	if rideCompletionResponse == nil {
		log.Fatalf("empty response in completeRide\n")
	}

	if rideCompletionResponse.Success == false {
		log.Fatalf("could not complete ride ?? ")
	}
	fmt.Printf("got response: %v, status: %v, type %T\n", rideCompletionResponse, rideCompletionResponse.Success, rideCompletionResponse.Success)

}
