package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"

	rscd "myuber/internal/client/driver" // ride sharing client; driver type
	rspb "myuber/pkg/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

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
			rscd.AcceptRide(request.RiderId, driverID, client)
		}
	}

}
