#!/bin/bash

# Constants for x and y
x=1.5
y=2.5

echo "Running KNN calculation with x=$x, y=$y"

# Build protofiles
echo "Building protofiles..."
make build_proto

# Build Go binaries
echo "Building Go binaries..."
make build

# Run the servers
echo "Running servers..."
make run_servers SERVERS=7

# Function to run client with a specific k value
run_client() {
    local k=$1
    echo "Running client with k=$k"
    echo "Time taken for client to get result (k=$k):"
    (time go run cmd/client/main.go -x $x -y $y -k $k) 2>&1 | tee "${k}_50000_benchmark.txt"
    echo "----------------------------------------"
}

run_client 3
run_client 10
run_client 25
run_client 50
run_client 100
run_client 500

# Clean up
echo "Cleaning up..."
make clean

echo "All operations completed."
