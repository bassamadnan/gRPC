package main

import (
	"context"
	"fmt"
	"io"
	"log"

	rspb "myuber/pkg/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	BASE_SERVER_ADDR := "localhost:5050"
	driverID := "driver1"
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
		fmt.Printf("recieved request: %v\n", request)

	}

}
