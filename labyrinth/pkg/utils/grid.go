package utils

import (
	lrpb "labyrinth/pkg/proto"
)

func CheckBounds(direction lrpb.Move_Direction, currX, currY, M, N int) bool {
	switch direction {
	case lrpb.Move_UP:
		return currY > 0
	case lrpb.Move_DOWN:
		return currY < M-1
	case lrpb.Move_LEFT:
		return currX > 0
	case lrpb.Move_RIGHT:
		return currX < N-1
	default:
		return false
	}
}

func GetNewPosition(direction lrpb.Move_Direction, currX, currY int) (int, int) {
	switch direction {
	case lrpb.Move_UP:
		return currX, currY - 1
	case lrpb.Move_DOWN:
		return currX, currY + 1
	case lrpb.Move_LEFT:
		return currX - 1, currY
	case lrpb.Move_RIGHT:
		return currX + 1, currY
	default:
		return currX, currY
	}
}

func CheckTile(direction lrpb.Move_Direction, currX, currY int, grid [][]string, tileType string) bool {
	newX, newY := GetNewPosition(direction, currX, currY)
	if newY >= 0 && newY < len(grid) && newX >= 0 && newX < len(grid[0]) {
		return grid[newY][newX] == tileType
	}
	return false
}

func UpdatePlayerPosition(direction lrpb.Move_Direction, currX, currY int, grid [][]string) (int, int, bool) {
	newX, newY := GetNewPosition(direction, currX, currY)
	coinCollected := false
	if grid[newY][newX] == "C" {
		grid[newY][newX] = "E"
		coinCollected = true
	}
	return newX, newY, coinCollected
}
