package main

import (
	"fmt"
	"math/rand"
	"strings"
)

const (
	W = "W" // Wall
	C = "C" // Coin
	G = "G" // Goal
	P = "P" // Player
	E = "E" // Empty
)

func main() {
	M, N := 5, 9

	grid := generateRandomGrid(M, N)
	grid[0][0] = P
	grid[M-1][N-1] = G

	printCSV(grid, M, N)
}

func generateRandomGrid(M, N int) [][]string {

	grid := make([][]string, M)
	for i := range grid {
		grid[i] = make([]string, N)
		for j := range grid[i] {
			grid[i][j] = getRandomCell()
		}
	}
	return grid
}

func getRandomCell() string {
	options := []string{W, C, E}
	return options[rand.Intn(len(options))]
}

func printCSV(grid [][]string, M, N int) {
	fmt.Printf("%d,%d\n", M, N)

	for _, row := range grid {
		fmt.Println(strings.Join(row, ","))
	}
}
