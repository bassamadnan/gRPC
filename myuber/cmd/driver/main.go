package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"

	auth "myuber/internal/auth"
	rscd "myuber/internal/client/driver" // ride sharing client; driver type
	utils "myuber/internal/utils"
	rspb "myuber/pkg/proto"

	"google.golang.org/grpc"
)

func main() {
	BASE_SERVER_ADDR := "localhost:5050"
	AVAILABLE := true
	clientId := flag.Int("id", 1, "driver id")
	flag.Parse()
	driverID := fmt.Sprintf("driver%v", *clientId)

	tlsCredentials, err := auth.ClientLoadTLSCredentials()
	if err != nil {
		log.Fatal("cannot load TLS credentials: ", err)
	}
	opts := []grpc.DialOption{grpc.WithTransportCredentials(tlsCredentials)}
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
		fmt.Printf("Driver: %v listening for ride requests\n", driverID)
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
		input := utils.TakeInput("accept/reject")

		if input == "1" {
			AVAILABLE = false
			resp := rscd.AcceptRide(request.RiderId, driverID, client)
			if resp.Success == false {
				fmt.Printf("rider: %v not waiting anymore\n", request.RiderId)
				continue
			}
			input := utils.TakeInput("complete")
			if input == "1" {
				rscd.CompleteRide(driverID, client)
			}
			fmt.Printf("Driver %v now available\n", AVAILABLE)
			AVAILABLE = true
		}
		if input == "0" {
			rscd.RejectRide(request.RiderId, driverID, client)
		}

	}

}
