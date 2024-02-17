package main

import (
	"github.com/vicdevcode/drivee-test/internal/alg"
)

func main() {
	clusters := alg.SetClusters()
	for i := range clusters {
		alg.HeldKapr(clusters[i].Cluster)
	}
}
