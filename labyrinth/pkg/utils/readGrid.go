package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

func LoadGridFromFile(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var grid [][]string
	scanner := bufio.NewScanner(file)

	if !scanner.Scan() {
		return nil, fmt.Errorf("failed to read dimensions")
	}

	for scanner.Scan() {
		line := scanner.Text()
		row := strings.Split(line, ",")
		grid = append(grid, row)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return grid, nil
}

func PrintGridAsTable(grid [][]string) {
	t := table.NewWriter()

	for _, row := range grid {
		tableRow := make(table.Row, len(row))
		for j, cell := range row {
			tableRow[j] = colorizeCell(cell)
		}
		t.AppendRow(tableRow)
	}

	t.SetStyle(table.StyleBold)
	fmt.Println(t.Render())
}

func colorizeCell(cell string) string {
	switch cell {
	case "W":
		return text.FgRed.Sprintf("W") // wall (red)
	case "C":
		return text.FgHiYellow.Sprintf("C") // coin (gold)
	default:
		return cell // Empty (no color)
	}
}
