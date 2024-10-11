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

func main() {
	lis, err := net.Listen("tcp", LOAD_BALANCER_ADDR)
	numServers := flag.Int("n", 1, "number of servers")
	algo := flag.Int("t", 1, "algorithm 1/2/3 fcfs/rr/custom")
	flag.Parse()
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// 5050
	// 5051
	// 5052.. assume this isnt more than 9
	weightedServers := make([]weightedServer, 0, *numServers)
	totalWeight := 0
	for i := 0; i < *numServers; i++ {
		addr := fmt.Sprintf("localhost:505%v", i)
		weight := rand.Intn(5) + 1 // random weight between 1 and 5
		weightedServers = append(weightedServers, weightedServer{address: addr, weight: weight})
		totalWeight += weight
		println(addr)
	}
	s := grpc.NewServer()
	lbpb.RegisterLoadBalanceServiceServer(s, &server{
		servers:     weightedServers,
		algorithm:   *algo,
		index:       0,
		totalWeight: totalWeight,
	})
	log.Printf("State server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
