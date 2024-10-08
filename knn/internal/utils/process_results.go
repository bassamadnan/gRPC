package utils

import (
	knnpb "knn/pkg/api"
	"math"
)

func ProcessResults(k int, results chan []*knnpb.Points) []*knnpb.Points {
	serverResults := make([][]*knnpb.Points, 0)
	for result := range results {
		serverResults = append(serverResults, result)
	}
	numServers := len(serverResults)

	indexes := make([]int, numServers)

	kNearest := make([]*knnpb.Points, 0, k)

	for len(kNearest) < k {
		minDistance := math.Inf(1)
		minServer := -1

		for i := 0; i < numServers; i++ {
			if indexes[i] < len(serverResults[i]) {
				distance := serverResults[i][indexes[i]].Distance
				if distance < minDistance {
					minDistance = distance
					minServer = i
				}
			}
		}

		if minServer == -1 {
			break
		}

		kNearest = append(kNearest, serverResults[minServer][indexes[minServer]])

		indexes[minServer]++
	}

	// fmt.Printf("indexes values: %v\n", indexes)
	return kNearest
}
