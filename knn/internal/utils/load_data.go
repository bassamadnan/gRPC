package utils

import (
	"bufio"
	"fmt"
	"knn/internal/knn"
	"os"
	"strconv"
	"strings"
)

func ReadFile(path string, start, end int) ([]knn.Point, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	var points []knn.Point
	scanner := bufio.NewScanner(file)

	lineNum := 0
	for scanner.Scan() {
		lineNum++
		if lineNum < start {
			continue
		}
		if lineNum > end && end != -1 {
			break
		}

		line := scanner.Text()
		fields := strings.Fields(line)

		if len(fields) != 2 {
			return nil, fmt.Errorf("invalid line format: %s", line)
		}

		x, err := strconv.ParseFloat(fields[0], 64)
		if err != nil {
			return nil, fmt.Errorf("error parsing X coordinate: %v", err)
		}

		y, err := strconv.ParseFloat(fields[1], 64)
		if err != nil {
			return nil, fmt.Errorf("error parsing Y coordinate: %v", err)
		}

		points = append(points, knn.Point{X: x, Y: y})
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	return points, nil
}

func CountFileLines(path string) (int, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineCount := 0
	for scanner.Scan() {
		lineCount++
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("error reading file: %v", err)
	}

	return lineCount, nil
}
