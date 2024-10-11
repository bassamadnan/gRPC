package main

import (
	"context"
	"fmt"
	"log"
	auth "myuber/internal/auth"
	interceptor "myuber/internal/interceptor"
	statemgmt "myuber/internal/server"
	rspb "myuber/pkg/proto/rideshare"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"
)

type server struct {
	rspb.UnimplementedRideServiceServer
	state            []statemgmt.RideInfo                                          // stores all the ride information
	streams          map[string]grpc.ServerStreamingServer[rspb.DriverRideRequest] // driverid, stream
	mu               sync.Mutex
	connectedDrivers []string
	timeout          int // timeout for int seconds
}

func (s *server) ConnectDriver(req *rspb.DriverRequest, stream grpc.ServerStreamingServer[rspb.DriverRideRequest]) error {
	s.mu.Lock()
	s.streams[req.DriverId] = stream
	s.connectedDrivers = append(s.connectedDrivers, req.DriverId)
	s.mu.Unlock()

	// https://pkg.go.dev/google.golang.org/grpc#ServerStream
	// keep alive
	// wait until stream is closed, returns a channel that closes when time out or context is cancelled?
	// related to Go in general, !!!
	<-stream.Context().Done()
	return nil
}

func (s *server) GetAvailableDrivers() []string {
	busyDrivers := make(map[string]bool)
	for _, ride := range s.state {
		if ride.DriverId != "" {
			busyDrivers[ride.DriverId] = true
		}
	}

	availableDrivers := []string{}
	for _, driverID := range s.connectedDrivers {
		if !busyDrivers[driverID] {
			availableDrivers = append(availableDrivers, driverID)
		}
	}

	return availableDrivers
}

func (s *server) isDriverAssigned(driverId string) bool {
	// is driver assigned to a rider?
	// should be present in the state
	s.mu.Lock()
	defer s.mu.Unlock()
	// check if assigned
	for _, rides := range s.state {
		if rides.DriverId == driverId {
			return true
		}
	}
	return false
}

func (s *server) isRiderAssigned(riderId string) bool {
	// is rider assigned to a driver? check by state
	// returns true if rider isnt in queue (no drivers available or completed or in progress already returned)
	s.mu.Lock()
	defer s.mu.Unlock()
	// check if assigned
	found := false
	for i, rides := range s.state {
		if rides.RiderId != riderId {
			continue
		}
		if s.state[i].Status != rspb.RideStatusResponse_PENDING {
			found = true
			break // rider isnt waiting for a driver, break the loop
		}
	}
	return found
}

func (s *server) AssignDriver(req *rspb.RideRequest) {
	availableDrivers := s.GetAvailableDrivers()
	stream_req := &rspb.DriverRideRequest{
		RiderId:     req.RiderId,
		Pickup:      req.Pickup,
		Destination: req.Destination,
	}
	for _, DriverId := range availableDrivers {
		if s.isDriverAssigned(DriverId) {
			continue
		}
		stream := s.streams[DriverId]
		err := stream.Send(stream_req)
		if err != nil {
			log.Fatalf("stream send error: %v\n", err)
		}
		// sleep this go routine for timeout seconds before sending next request
		fmt.Printf("sent rider %v's request to driver %v\n", req.RiderId, DriverId)
		time.Sleep(time.Duration(s.timeout) * time.Second)
		if s.isRiderAssigned(req.RiderId) {
			return
		}
	}

	if !s.isRiderAssigned(req.RiderId) {
		for i, rides := range s.state {
			if rides.RiderId != req.RiderId {
				continue
			}
			s.state[i].Status = rspb.RideStatusResponse_NO_DRIVERS_AVAILABLE // none could be assigned
		}
	}
}

func (s *server) RequestRide(ctx context.Context, req *rspb.RideRequest) (*rspb.RideResponse, error) {
	s.mu.Lock()
	for i, rides := range s.state {
		// if entry already exists, reset status to pending
		if rides.RiderId != req.RiderId {
			continue
		}
		s.state[i].Status = rspb.RideStatusResponse_PENDING
	}

	// create new rideinfo instance
	currRideInfo := &statemgmt.RideInfo{
		RiderId:     req.RiderId,
		Destination: req.Destination,
		Pickup:      req.Pickup,
		Status:      rspb.RideStatusResponse_PENDING, // finding a driver
	}
	availableDrivers := s.GetAvailableDrivers()
	if len(availableDrivers) == 0 {
		fmt.Printf("No rides available\n")
		currRideInfo.Status = rspb.RideStatusResponse_NO_DRIVERS_AVAILABLE
		s.state = append(s.state, *currRideInfo)
		return &rspb.RideResponse{
			Success: false,
		}, nil
	}

	s.state = append(s.state, *currRideInfo)
	s.mu.Unlock()
	go s.AssignDriver(req) // go routine to assign drivers to be run in background

	return &rspb.RideResponse{
		Success: true, // attempting to find a driver, request placed sucessfully
	}, nil
}

func (s *server) GetRideStatus(c context.Context, req *rspb.RideStatusRequest) (*rspb.RideStatusResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, rides := range s.state {
		if rides.RiderId != req.RiderId {
			continue
		}
		return &rspb.RideStatusResponse{
			RiderId:  rides.RiderId,
			DriverId: rides.DriverId,
			Status:   rides.Status,
		}, nil
	}
	return nil, nil
}

func (s *server) AcceptRide(ctx context.Context, req *rspb.AcceptRideRequest) (*rspb.AcceptRideResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i, rides := range s.state {
		if rides.RiderId != req.RiderId {
			continue
		}

		if rides.Status == rspb.RideStatusResponse_PENDING {
			// successfully assign
			s.state[i].DriverId = req.DriverId
			s.state[i].Status = rspb.RideStatusResponse_IN_PROGRESS
			return &rspb.AcceptRideResponse{
				Success: true,
			}, nil
		}

		// already accepted by some other driver (in progress)
		// or client alreayd got a no driver available response, client expected to request again
		return &rspb.AcceptRideResponse{
			Success: false,
		}, nil
	}
	return nil, nil
}

func (s *server) CompleteRide(ctx context.Context, req *rspb.RideCompletionRequest) (*rspb.RideCompletionResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i, rides := range s.state {
		if rides.DriverId != req.DriverId {
			continue
		}
		if rides.Status != rspb.RideStatusResponse_IN_PROGRESS {
			// should never reach here
			return &rspb.RideCompletionResponse{
				Success: false,
			}, nil
		}

		s.state[i].Status = rspb.RideStatusResponse_COMPLETED
		s.state[i].DriverId = ""
		return &rspb.RideCompletionResponse{
			Success: true,
		}, nil
	}
	return nil, nil
}

func main() {
	port := 5050
	DEFAULT_TIMEOUT := 5
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	tlsCredentials, err := auth.ServerLoadTLSCredentials()
	if err != nil {
		log.Fatal("cannot load TLS credentials: ", err)
	}
	var opts []grpc.ServerOption
	chainedInterceptor := interceptor.ChainedInterceptor(
		interceptor.AuthorizationInterceptor,
		interceptor.LoggingInterceptor,
	)
	opts = []grpc.ServerOption{grpc.Creds(tlsCredentials), grpc.UnaryInterceptor(chainedInterceptor)}
	s := grpc.NewServer(opts...)
	rspb.RegisterRideServiceServer(s, &server{
		streams:          make(map[string]grpc.ServerStreamingServer[rspb.DriverRideRequest]),
		state:            make([]statemgmt.RideInfo, 0),
		connectedDrivers: make([]string, 0),
		timeout:          DEFAULT_TIMEOUT,
	})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
