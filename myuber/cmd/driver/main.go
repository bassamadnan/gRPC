package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"time"

	rspb "myuber/pkg/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func acceptRide(riderId string, driverId string, c rspb.RideServiceClient) {
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
}

func rejectRide(riderId string, driverId string, c rspb.RideServiceClient) {
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

func main() {
	BASE_SERVER_ADDR := "localhost:5050"
	clientId := flag.Int("id", 1, "driver id")
	flag.Parse()
	driverID := fmt.Sprintf("driver%v", *clientId)
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	conn, err := grpc.NewClient(BASE_SERVER_ADDR, opts...)
	if err != nil {
		log.Fatalf("conn failed %v", err)
	}
	defer conn.Close()

	client := rspb.NewRideServiceClient(conn)
	// open a stream for recieving ride requests from server
	stream, err := client.ConnectDriver(context.Background(), &rspb.DriverRequest{DriverId: driverID})
	if err != nil {
		log.Fatalf("stream error %v\n", err)
	}

	for {
		request, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("Stream ended")
			break
		}

		if err != nil {
			log.Fatalf("error recv stream %v\n", err)
		}
		// take user input for accept/reject
		fmt.Printf("recieved request: %v, %T\n", request, request)
		var input string
		fmt.Print("Enter 1/0 to accept/reject: ")
		n, err := fmt.Scanf("%v", &input)
		if err != nil || n != 1 {
			log.Fatalf("scanf error %v", err)
		}

		if input == "1" {
			acceptRide(request.RiderId, driverID, client)
		}
	}

}
