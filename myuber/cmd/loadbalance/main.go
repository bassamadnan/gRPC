package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/rand"
	lbpb "myuber/pkg/proto/loadbalance"
	"net"
	"sync"

	"google.golang.org/grpc"
)

// GetServers(context.Context, *Empty) (*Servers, error)
// GetServer(context.Context, *Empty) (*Server, error)
// AddServer(context.Context, *Server) (*Empty, error)

const (
	LOAD_BALANCER_ADDR string = "localhost:7070"
	FIRST_PICK                = 1 // pick at random
	ROUND_ROBIN               = 2 // increment index at each call
	CUSTOM                    = 3 // weighted round robin
)

type weightedServer struct {
	address string
	weight  int
}

type server struct {
	lbpb.UnimplementedLoadBalanceServiceServer
	servers     []weightedServer // all servers known to load balancer
	algorithm   int              // load balancing algorithm
	index       int              // for round robin
	totalWeight int              // total weight of all servers
	mu          sync.Mutex
}

// acts as a client side service discovery

func (s *server) GetServers(ctx context.Context, req *lbpb.Empty) (*lbpb.Servers, error) {
	addresses := make([]string, len(s.servers))
	for i, ws := range s.servers {
		addresses[i] = ws.address
	}
	return &lbpb.Servers{
		Servers: addresses,
	}, nil
}

func (s *server) GetServer(ctx context.Context, req *lbpb.Empty) (*lbpb.Server, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if len(s.servers) == 0 {
		return nil, fmt.Errorf("no servers available")
	}

	var selectedServer string

	switch s.algorithm {
	case FIRST_PICK:
		selectedServer = s.servers[rand.Intn(len(s.servers))].address
	case ROUND_ROBIN:
		selectedServer = s.servers[s.index].address
		s.index = (s.index + 1) % len(s.servers)
	case CUSTOM:
		randomWeight := rand.Intn(s.totalWeight) + 1 // sample a number
		for _, ws := range s.servers {
			randomWeight -= ws.weight
			if randomWeight <= 0 {
				selectedServer = ws.address
				break
			}
		}
	default:
		selectedServer = s.servers[0].address
	}

	return &lbpb.Server{
		Server: selectedServer,
	}, nil
}

func (s *server) AddServer(ctx context.Context, req *lbpb.Server) (*lbpb.Empty, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, ws := range s.servers {
		if ws.address == req.Server {
			// should not come here
			return &lbpb.Empty{}, fmt.Errorf("server %s already exists", req.Server)
		}
	}
	// assign a weight
	weight := rand.Intn(5) + 1 // random weight between 1 and 5
	s.servers = append(s.servers, weightedServer{address: req.Server, weight: weight})
	s.totalWeight += weight

	log.Printf("Added new server: %s with weight %d", req.Server, weight)
	return &lbpb.Empty{}, nil
}

func main() {
	lis, err := net.Listen("tcp", LOAD_BALANCER_ADDR)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	algo := flag.Int("t", 1, "algorithm 1/2/3 fcfs/rr/custom")
	flag.Parse()

	s := grpc.NewServer()
	lbpb.RegisterLoadBalanceServiceServer(s, &server{
		servers:     []weightedServer{},
		algorithm:   *algo,
		index:       0,
		totalWeight: 0,
	})
	log.Printf("Load balancer listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
