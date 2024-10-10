package main

import (
	"context"
	"fmt"
	"log"
	rspb "myuber/pkg/proto"
	"net"
	"sync"

	"google.golang.org/grpc"
)

// var (
//
//	tls        = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
//	certFile   = flag.String("cert_file", "", "The TLS cert file")
//	keyFile    = flag.String("key_file", "", "The TLS key file")
//	jsonDBFile = flag.String("json_db_file", "", "A json file containing a list of features")
//	port       = flag.Int("port", 50051, "The server port")
//
// )

// type RideServiceServer interface {
// 	RequestRide(context.Context, *RideRequest) (*RideResponse, error)
// 	GetRideStatus(context.Context, *RideStatusRequest) (*RideStatusResponse, error)
// 	ConnectDriver(*DriverRequest, grpc.ServerStreamingServer[DriverRideRequest]) error
// 	AcceptRide(context.Context, *AcceptRideRequest) (*AcceptRideResponse, error)
// 	RejectRide(context.Context, *RejectRideRequest) (*RejectRideResponse, error)
// 	CompleteRide(context.Context, *RideCompletionRequest) (*RideCompletionResponse, error)
// }

type rideInfo struct {
	rider_id    string
	driver_id   string
	pickup      string
	destination string
	status      rspb.RideStatusResponse_Status
}

type server struct {
	rspb.UnimplementedRideServiceServer
	state            []rideInfo                                                    // stores all the ride information
	streams          map[string]grpc.ServerStreamingServer[rspb.DriverRideRequest] // driverid, stream
	mu               sync.Mutex
	connectedDrivers []string
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
		if ride.driver_id != "" {
			busyDrivers[ride.driver_id] = true
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
func (s *server) AcceptRide(ctx context.Context, req *rspb.AcceptRideRequest) (*rspb.AcceptRideResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i, rides := range s.state {
		if rides.rider_id != req.RiderId {
			continue
		}

		if rides.status == rspb.RideStatusResponse_PENDING {
			// successfully assign
			s.state[i].driver_id = req.DriverId
			s.state[i].status = rspb.RideStatusResponse_IN_PROGRESS
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

func (s *server) RequestRide(ctx context.Context, req *rspb.RideRequest) (*rspb.RideResponse, error) {
	fmt.Println("inside req ride ")
	// iterate over all drivers
	// for each driver send a request, with time out
	// if no driver found/all occupied /all timed out then just return
	// if driver is found then
	// send requests to all drivers
	// assign the driver
	s.mu.Lock()
	defer s.mu.Unlock()

	// create new rideinfo instance
	currRideInfo := &rideInfo{
		rider_id:    req.RiderId,
		destination: req.Destination,
		pickup:      req.Pickup,
		status:      rspb.RideStatusResponse_PENDING, // finding a driver
	}
	availableDrivers := s.GetAvailableDrivers()
	if len(availableDrivers) == 0 {
		fmt.Printf("No rides available\n")
		currRideInfo.status = rspb.RideStatusResponse_NO_DRIVERS_AVAILABLE
		s.state = append(s.state, *currRideInfo)
		return &rspb.RideResponse{
			Success: false,
		}, nil
	}

	s.state = append(s.state, *currRideInfo)
	// DriverRideRequest ->  to be sent via streams to drivers
	stream_req := &rspb.DriverRideRequest{
		RiderId:     req.RiderId,
		Pickup:      req.Pickup,
		Destination: req.Destination,
	}
	for _, DriverId := range availableDrivers {
		stream := s.streams[DriverId]
		err := stream.Send(stream_req)
		if err != nil {
			log.Fatalf("stream send error: %v\n", err)
		}
	}
	return &rspb.RideResponse{
		Success: true, // attempting to find a driver, request placed sucessfully
	}, nil
}
func (s *server) GetRideStatus(c context.Context, req *rspb.RideStatusRequest) (*rspb.RideStatusResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, rides := range s.state {
		if rides.rider_id != req.RiderId {
			continue
		}
		return &rspb.RideStatusResponse{
			RiderId:  rides.rider_id,
			DriverId: rides.driver_id,
			Status:   rides.status,
		}, nil
	}
	return nil, nil
}

func unaryInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	log.Println("--> unary interceptor: ", info.FullMethod)
	return handler(ctx, req)
}

func main() {
	port := 5050
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer(grpc.UnaryInterceptor(unaryInterceptor))
	rspb.RegisterRideServiceServer(s, &server{
		streams:          make(map[string]grpc.ServerStreamingServer[rspb.DriverRideRequest]),
		state:            make([]rideInfo, 0),
		connectedDrivers: make([]string, 0),
	})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
