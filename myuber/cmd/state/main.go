package main

import (
	"context"
	"log"
	spb "myuber/pkg/proto/state"
	"net"
	"sync"

	"google.golang.org/grpc"
)

const BASE_SERVER_ADDR string = "localhost:6969"

type rideInfo struct {
	rider_id    string
	driver_id   string
	pickup      string
	destination string
	status      spb.RideInfo_Status
}

type server struct {
	spb.UnimplementedStateServiceServer
	state  []rideInfo
	mu     sync.Mutex
	holder string // server holding the lock
}

// type StateServiceServer interface {
// 	GetState(context.Context, *StateRequest) (*State, error)
// 	SetState(context.Context, *State) (*StateResponse, error)
// 	mustEmbedUnimplementedStateServiceServer()
// }

func (s *server) GetState(ctx context.Context, req *spb.StateRequest) (*spb.State, error) {
	s.mu.Lock()
	state_arr := make([]*spb.RideInfo, 0, len(s.state))
	for _, ride := range s.state {
		rInfo := &spb.RideInfo{
			RiderId:     ride.rider_id,
			DriverId:    ride.driver_id,
			Destination: ride.destination,
			Pickup:      ride.pickup,
			Status:      ride.status,
		}
		state_arr = append(state_arr, rInfo)
	}
	s.holder = req.Server
	log.Printf("lock given to server %v --> %v\n", req.Server, s.state)
	return &spb.State{
		State: state_arr,
	}, nil
}

func (s *server) SetState(ctx context.Context, req *spb.State) (*spb.StateResponse, error) {
	incoming_server := req.Server
	if incoming_server != s.holder {
		log.Fatalf("Invalid lock release request from %v , current holder: %v\n", incoming_server, s.holder)
		return &spb.StateResponse{
			Success: false,
		}, nil
	}
	defer s.mu.Unlock()
	s.state = make([]rideInfo, 0, len(req.State))
	for _, rideInfoProto := range req.State {
		ride := rideInfo{
			rider_id:    rideInfoProto.RiderId,
			driver_id:   rideInfoProto.DriverId,
			pickup:      rideInfoProto.Pickup,
			destination: rideInfoProto.Destination,
			status:      rideInfoProto.Status,
		}
		s.state = append(s.state, ride)
	}
	log.Printf("Lock released by %s\n", s.holder)
	s.holder = ""
	log.Printf("state after setState %v\n", s.state)
	return &spb.StateResponse{
		Success: true,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", BASE_SERVER_ADDR)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	spb.RegisterStateServiceServer(s, &server{
		state: make([]rideInfo, 0),
	})
	log.Printf("State server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
