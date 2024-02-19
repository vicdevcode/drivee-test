package main

import (
	"fmt"

	"github.com/vicdevcode/drivee-test/internal/alg"
)

func main() {
	clusters := alg.SetClusters()
	for _, cluster := range clusters {
		if len(cluster.Cluster) > 2 {
			fmt.Println(alg.SolveGA(len(cluster.Cluster), 10, cluster.Cluster))
		} else {
			fmt.Println(cluster)
		}
	}
}
