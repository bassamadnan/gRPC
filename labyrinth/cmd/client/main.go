package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

// Table characters
const (
	H = "H" // Hidden, no color
	W = "W" // Wall, red color
	C = "C" // Coin, gold color
	G = "G" // Goal, blue color
	P = "P" // Player, yellow or choice color
	E = " " // Empty
)

func main() {
	// Static size for now, but can be taken as input
	M, N := 18, 18

	// Create the grid
	grid := generateRandomGrid(M, N)

	// Ensure P (player) at top-left and G (goal) at bottom-right
	grid[0][0] = P
	grid[M-1][N-1] = G

	// Print both tables side by side: grid and player status
	printTablesSideBySide(grid, M, N)
}

// generateRandomGrid creates an MxN grid with random characters (H, W, C, E)
func generateRandomGrid(M, N int) [][]string {
	rand.Seed(time.Now().UnixNano())

	grid := make([][]string, M)
	for i := range grid {
		grid[i] = make([]string, N)
		for j := range grid[i] {
			// Randomly assign H, W, C, or E (empty)
			grid[i][j] = getRandomCell()
		}
	}
	return grid
}

// getRandomCell returns a random cell character (H, W, C, or E)
func getRandomCell() string {
	options := []string{H, W, C, E}
	return options[rand.Intn(len(options))]
}

// printTablesSideBySide prints both the grid and the status table side by side
func printTablesSideBySide(grid [][]string, M, N int) {
	// Create the grid table
	tGrid := table.NewWriter()

	for _, row := range grid {
		tableRow := make(table.Row, len(row))
		for j, cell := range row {
			paddedCell := "  " + colorizeCell(cell) + "  "
			tableRow[j] = paddedCell
		}
		tGrid.AppendRow(tableRow)
	}

	// Customize table style for the grid
	tGrid.SetStyle(table.Style{
		Box: table.StyleBoxBold,
		Color: table.ColorOptions{
			Header: text.Colors{text.FgHiWhite},
			Row:    text.Colors{text.Reset},
		},
		Options: table.Options{
			DrawBorder:      true,
			SeparateColumns: true,
			SeparateRows:    true,
		},
	})

	// Create the status table
	tStatus := table.NewWriter()
	health := rand.Intn(100) // Random health between 0-100
	x := rand.Intn(M)        // Random X between 0 and M-1
	y := rand.Intn(N)        // Random Y between 0 and N-1
	move := "UP"             // Default move
	status := "SUCCESS"      // Status is always "SUCCESS"

	// Append status rows
	tStatus.AppendRow(table.Row{"Health", health})
	tStatus.AppendRow(table.Row{"X", x})
	tStatus.AppendRow(table.Row{"Y", y})
	tStatus.AppendRow(table.Row{"Move", move}) // Added Move row
	tStatus.AppendRow(table.Row{"Status", status})

	// Customize the style for the status table
	tStatus.SetStyle(table.Style{
		Box: table.StyleBoxLight,
		Color: table.ColorOptions{
			Row: text.Colors{text.Reset},
		},
		Options: table.Options{
			DrawBorder:      true,
			SeparateColumns: true,
		},
	})

	// Render both tables to string slices (line by line)
	gridLines := strings.Split(tGrid.Render(), "\n")
	statusLines := strings.Split(tStatus.Render(), "\n")

	// Ensure both tables have the same number of lines
	maxLines := int(math.Max(float64(len(gridLines)), float64(len(statusLines))))

	// Append empty lines to the shorter table
	for len(gridLines) < maxLines {
		gridLines = append(gridLines, "")
	}
	for len(statusLines) < maxLines {
		statusLines = append(statusLines, "")
	}

	// Print both tables side by side
	for i := 0; i < maxLines; i++ {
		fmt.Printf("%-80s %s\n", gridLines[i], statusLines[i])
	}
}

// colorizeCell adds color to the cell based on its character
func colorizeCell(cell string) string {
	switch cell {
	case W:
		return text.FgRed.Sprintf(W) // Wall (red)
	case C:
		return text.FgHiYellow.Sprintf(C) // Coin (gold)
	case G:
		return text.FgBlue.Sprintf(G) // Goal (blue)
	case P:
		return text.FgYellow.Sprintf(P) // Player (yellow or chosen color)
	default:
		return cell // No color (hidden or empty)
	}
}
