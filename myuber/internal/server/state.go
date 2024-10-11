package state

import (
	"context"
	"fmt"
	"log"
	rspb "myuber/pkg/proto/rideshare"
	spb "myuber/pkg/proto/state"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const STATE_SERVICE_ADDR = "localhost:6969"
const STATE_TIMEOUT = 60

type RideInfo struct {
	RiderId     string
	DriverId    string
	Pickup      string
	Destination string
	Status      rspb.RideStatusResponse_Status
}

func connectToServer() (*grpc.ClientConn, spb.StateServiceClient, error) {
	conn, err := grpc.NewClient(STATE_SERVICE_ADDR, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect: %v", err)
	}
	client := spb.NewStateServiceClient(conn)
	return conn, client, nil
}

func getRspbStatus(status spb.RideInfo_Status) rspb.RideStatusResponse_Status {
	switch status {
	case spb.RideInfo_PENDING:
		return rspb.RideStatusResponse_PENDING
	case spb.RideInfo_IN_PROGRESS:
		return rspb.RideStatusResponse_IN_PROGRESS
	case spb.RideInfo_NO_DRIVERS_AVAILABLE:
		return rspb.RideStatusResponse_NO_DRIVERS_AVAILABLE
	case spb.RideInfo_COMPLETED:
		return rspb.RideStatusResponse_COMPLETED
	default:
		log.Printf("Unknown status: %v", status)
		return rspb.RideStatusResponse_PENDING
	}
}

func getSpbStatus(status rspb.RideStatusResponse_Status) spb.RideInfo_Status {
	switch status {
	case rspb.RideStatusResponse_PENDING:
		return spb.RideInfo_PENDING
	case rspb.RideStatusResponse_IN_PROGRESS:
		return spb.RideInfo_IN_PROGRESS
	case rspb.RideStatusResponse_NO_DRIVERS_AVAILABLE:
		return spb.RideInfo_NO_DRIVERS_AVAILABLE
	case rspb.RideStatusResponse_COMPLETED:
		return spb.RideInfo_COMPLETED
	default:
		log.Printf("Unknown status: %v", status)
		return spb.RideInfo_PENDING
	}
}

func GetState(serverName string) ([]RideInfo, error) {
	conn, client, err := connectToServer()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*STATE_TIMEOUT)
	defer cancel()

	req := &spb.StateRequest{Server: serverName}
	state, err := client.GetState(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("error in getting state: %v", err)
	}
	rideInfoArr := make([]RideInfo, 0, len(state.State))
	for _, ride := range state.State {
		rideInfoArr = append(rideInfoArr, RideInfo{
			RiderId:     ride.RiderId,
			DriverId:    ride.DriverId,
			Pickup:      ride.Pickup,
			Destination: ride.Destination,
			Status:      getRspbStatus(ride.Status),
		})
	}
	return rideInfoArr, nil
}

func SetState(state []RideInfo, serverName string) (*spb.StateResponse, error) {
	conn, client, err := connectToServer()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*STATE_TIMEOUT)
	defer cancel()

	spbState := &spb.State{
		State:  make([]*spb.RideInfo, len(state)),
		Server: serverName,
	}

	for i, ride := range state {
		spbState.State[i] = &spb.RideInfo{
			RiderId:     ride.RiderId,
			DriverId:    ride.DriverId,
			Pickup:      ride.Pickup,
			Destination: ride.Destination,
			Status:      getSpbStatus(ride.Status),
		}
	}

	return client.SetState(ctx, spbState)
}
