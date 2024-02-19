package main

import (
	"fmt"

	"github.com/vicdevcode/drivee-test/internal/alg"
)

func main() {
	orders, couriers := alg.FakeData()
	if len(couriers) == 1 {
		couriers[0].Orders = orders
		for i := range orders {
			orders[i].CourierID = &couriers[0].ID
		}
		cluster := alg.SetClusters(couriers)[0]
		pathTime, path := alg.SolveGA(len(cluster.Cluster), 20, cluster.Cluster)
		path = path[1:]
		for i := range path {
			path[i] = path[i] - 1
		}
		couriers[0].Path = path
		couriers[0].FullPathTime = pathTime
		fmt.Println(couriers[0])
		return
	}
	orders, couriers = alg.HungarianAlg(orders, couriers)
	clusters := alg.SetClusters(couriers)

	for _, cluster := range clusters {
		if len(cluster.Cluster) > 2 {
			pathTime, path := alg.SolveGA(len(cluster.Cluster), 20, cluster.Cluster)
			path = path[1:]
			for i := range path {
				path[i] = path[i] - 1
			}
			couriers[cluster.CourierID].Path = path
			couriers[cluster.CourierID].FullPathTime = pathTime
		} else {
			fmt.Println(cluster)
		}
	}
	for _, courier := range couriers {
		fmt.Println(courier)
	}
}
