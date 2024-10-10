package main

import (
	"context"
	"fmt"
	"log"
	rspb "myuber/pkg/proto"
	"net"

	"google.golang.org/grpc"
)

//	type RideServiceServer interface {
//		RequestRide(context.Context, *RideRequest) (*RideResponse, error)
//		GetRideStatus(context.Context, *RideStatusRequest) (*RideStatusResponse, error)
//		AcceptRide(context.Context, *AcceptRideRequest) (*AcceptRideResponse, error)
//		RejectRide(context.Context, *RejectRideRequest) (*RejectRideResponse, error)
//		CompleteRide(context.Context, *RideCompletionRequest) (*RideCompletionResponse, error)
//		mustEmbedUnimplementedRideServiceServer()
//	}

// var (
//
//	tls        = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
//	certFile   = flag.String("cert_file", "", "The TLS cert file")
//	keyFile    = flag.String("key_file", "", "The TLS key file")
//	jsonDBFile = flag.String("json_db_file", "", "A json file containing a list of features")
//	port       = flag.Int("port", 50051, "The server port")
//
// )

type rideInfo struct {
	rider_id       string
	driver_id      string
	response       rspb.RideResponse_Status
	statusResponse rspb.RideStatusResponse_Status
	pickup         string
	destination    string
}

func CreateNewRideInfo(rider_id string) rideInfo {
	return rideInfo{
		rider_id:  rider_id,
		driver_id: "-1", //  no driver initially
		status:    0,    // 0 for pending
	}
}

type server struct {
	rspb.UnimplementedRideServiceServer

	state []rideInfo // stores all the ride information
}

func (s *server) RequestRide(context.Context, *rspb.RideRequest) (*rspb.RideResponse, error) {
	fmt.Println("inside req ride ")
	// iterate over all drivers
	// for each driver send a request, with time out
	// if no driver found/all occupied /all timed out then just return
	// if driver is found then
	// send accept ride request from the driver
	// assign the driver

}
func (s *server) GetRideStatus(context.Context, *rspb.RideStatusRequest) (*rspb.RideStatusResponse, error) {
	// time.Sleep(11 * time.Second) -> results in timeout
	fake_response := &rspb.RideStatusResponse{
		Status:   rspb.RideStatusResponse_PENDING,
		RiderId:  "riderlol",
		DriverId: "driverlol",
	}
	return fake_response, nil
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
	rspb.RegisterRideServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
