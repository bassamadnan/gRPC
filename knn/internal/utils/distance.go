package utils

import (
	"knn/internal/knn"
	"math"
	"sort"
)

func GetKNN(k uint16, point knn.Point, points []knn.Point) []knn.Distances {
	distances := make([]knn.Distances, len(points))
	for i, p := range points {
		dist := calculateDistance(point, p)
		distances[i] = knn.Distances{Point: p, Distance: dist}
	}

	sort.Slice(distances, func(i, j int) bool {
		return distances[i].Distance < distances[j].Distance
	})

	if int(k) > len(distances) {
		k = uint16(len(distances))
	}
	return distances[:k]
}

func calculateDistance(p1, p2 knn.Point) float64 {
	dx := p1.X - p2.X
	dy := p1.Y - p2.Y
	return math.Sqrt(dx*dx + dy*dy)
}
