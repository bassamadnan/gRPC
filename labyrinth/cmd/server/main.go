package main

import (
	"fmt"
	utils "labyrinth/pkg/utils"
)

func main() {
	grid, err := utils.LoadGridFromFile("grid.txt")
	if err != nil {
		fmt.Println("Error loading grid:", err)
		return
	}

	utils.PrintGridAsTable(grid)
}
