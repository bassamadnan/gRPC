package main

import (
	"flag"
	"fmt"
	"log"

	auth "myuber/internal/auth"
	rscr "myuber/internal/client/rider" //rider sharing client: rider type
	rspb "myuber/pkg/proto"

	"google.golang.org/grpc"
)

func main() {

	BASE_SERVER_ADDR := "localhost:5050"
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
	print(client)
	rscr.RequestRide(info, client)

	for {

		var input string
		fmt.Print("Enter 1 to get ride/status: ")
		n, err := fmt.Scanf("%v", &input)
		if err != nil || n != 1 {
			log.Fatalf("scanf error %v", err)
		}

		if input == "1" {
			rscr.GetRideStatus(riderid, client)
		}
	}
}
