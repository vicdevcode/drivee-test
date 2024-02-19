package alg

import (
	"sort"

	"github.com/arthurkushman/go-hungarian"

	"github.com/vicdevcode/drivee-test/internal/entity"
)

// func OptimizeCouriers() entity.Courier {
//
// }

func SetClusters() []entity.Cluster {
	orders, couriers := FakeData()

	COMatrix := CreateCourierOrderMatrix(orders, couriers)
	deletedId := []int{}

	for a := 0; a < len(orders); a += len(couriers) {
		sqMatrix := ConvertToSquareMatrix(COMatrix)

		hung := hungarian.SolveMin(sqMatrix)
		for courierId, value := range hung {
			if courierId >= len(couriers) {
				continue
			}
			for orderId, dist := range value {
				if dist > 0.0 {
					id := orderId
					for i := range deletedId {
						if i == 0 {
							if deletedId[i] <= id {
								id++
							}
							continue
						}
						if deletedId[i] >= orderId {
							id++
						}
					}
					orders[orderId].CourierID = &courierId
					couriers[courierId].Orders = append(couriers[courierId].Orders, orders[id])
					COMatrix = RemoveVisitedOrder(COMatrix, orderId)
					deletedId = append(deletedId, id)
					sort.Ints(deletedId)
				}
			}
		}
	}

	// Кластеры для каждого курьера
	clusters := []entity.Cluster{}

	for _, courier := range couriers {

		// Формируем кластер для одного курьера
		cluster := [][]float64{}

		courierRow := []float64{0}
		for _, order := range courier.Orders {
			courierRow = append(courierRow, calculateTime(courier.Speed, calculateDistance(courier.Coordinates, order.Start, order.End, courier.Penalty)))
		}
		cluster = append(cluster, courierRow)

		// Заполняем двумерную матрицу
		for i, order := range courier.Orders {
			cityRow := []float64{calculateTime(courier.Speed, calculateDistance(courier.Coordinates, order.Start, order.End, courier.Penalty))}
			for j, orderToOrder := range courier.Orders {
				if j == i {
					cityRow = append(cityRow, 0)
					continue
				}
				cityRow = append(cityRow, calculateTime(courier.Speed, calculateDistance(
					order.End,
					orderToOrder.Start,
					orderToOrder.End,
					courier.Penalty,
				)))
			}
			cluster = append(cluster, cityRow)
		}
		clusters = append(clusters, entity.Cluster{CourierID: courier.ID, Cluster: cluster})
	}

	return clusters
}
