package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"os/exec"
	"strings"

	lrpb "labyrinth/pkg/proto"
	utils "labyrinth/pkg/utils"

	"github.com/eiannone/keyboard"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	H = "H" // Hidden, no color
	W = "W" // Wall, red color
	C = "C" // Coin, gold color
	G = "G" // Goal, blue color
	P = "P" // Player, yellow or choice color
	E = " " // Empty
)

func main() {
	BASE_SERVER_ADDR := "localhost:5050"
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	conn, err := grpc.NewClient(BASE_SERVER_ADDR, opts...)
	if err != nil {
		log.Fatalf("conn failed %v", err)
	}
	defer conn.Close()
	client := utils.Client{
		Client: lrpb.NewLabyrinthServiceClient(conn),
		Status: "-",
	}

	M, N := client.GetLabyrinthInfo()
	client.GetPlayerStatus()

	grid := generateHiddenGrid(M, N)
	grid[client.Y][client.X] = P
	grid[M-1][N-1] = G

	startGame(&client, grid)
	fmt.Printf("Game over- result : %v\n", client.Status)
}

func startGame(client *utils.Client, grid [][]string) {
	if err := keyboard.Open(); err != nil {
		log.Fatal(err)
	}
	defer keyboard.Close()

	for {
		grid[client.Y][client.X] = P
		printTablesSideBySide(grid, client)
		if client.Status == "VICTORY" || client.Status == "DEATH" {
			return
		}
		char, key, err := keyboard.GetKey()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("You pressed: %c\n", char)

		if key == keyboard.KeyEsc || char == 'q' {
			break
		}

		switch char {
		case 'w', 'a', 's', 'd':
			client.HandleMove(char, grid)
		case '1':
			var x, y int
			var spellType string
			fmt.Print("Enter x y type: ")
			fmt.Scanf("%d %d %s", &x, &y, &spellType)
			fmt.Printf("Entered x :%v,  y: %v,  type: %v \n ", x, y, spellType)
		case '2':
			var x, y int
			fmt.Print("Enter x y: ")
			fmt.Scanf("%d %d", &x, &y)
			fmt.Printf("Entered x :%v,  y: %v\n ", x, y)
		}

		client.GetPlayerStatus()
	}
}

func generateHiddenGrid(M, N int) [][]string {
	grid := make([][]string, M)
	for i := range grid {
		grid[i] = make([]string, N)
		for j := range grid[i] {
			grid[i][j] = H
		}
	}
	return grid
}

func printTablesSideBySide(grid [][]string, client *utils.Client) {
	cmd := exec.Command("clear") //Linux example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()
	for i := 0; i < 15; i++ {
		fmt.Println()
	}
	tGrid := table.NewWriter()

	for _, row := range grid {
		tableRow := make(table.Row, len(row))
		for j, cell := range row {
			paddedCell := "  " + colorizeCell(cell) + "  "
			tableRow[j] = paddedCell
		}
		tGrid.AppendRow(tableRow)
	}
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

	tStatus := table.NewWriter()

	tStatus.AppendRow(table.Row{"Health", client.Health})
	tStatus.AppendRow(table.Row{"X", client.X})
	tStatus.AppendRow(table.Row{"Y", client.Y})
	tStatus.AppendRow(table.Row{"Score", client.Score})
	tStatus.AppendRow(table.Row{"Spells", client.Spells})
	tStatus.AppendRow(table.Row{"Status", client.Status})

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

	gridLines := strings.Split(tGrid.Render(), "\n")
	statusLines := strings.Split(tStatus.Render(), "\n")

	maxLines := int(math.Max(float64(len(gridLines)), float64(len(statusLines))))
	for len(gridLines) < maxLines {
		gridLines = append(gridLines, "")
	}
	for len(statusLines) < maxLines {
		statusLines = append(statusLines, "")
	}
	for i := 0; i < maxLines; i++ {
		fmt.Printf("%-80s %s\n", gridLines[i], statusLines[i])
	}
	for i := 0; i < 15; i++ {
		fmt.Println()
	}
}

func colorizeCell(cell string) string {
	switch cell {
	case H:
		return text.FgHiBlack.Sprintf(H) // Hidden (dark gray)
	case P:
		return text.FgYellow.Sprintf(P) // Player (yellow)
	case G:
		return text.FgBlue.Sprintf(G) // Goal (blue)
	default:
		return cell
	}
}
