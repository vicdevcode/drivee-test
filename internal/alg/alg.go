package alg

import (
	"fmt"
	"sort"

	"github.com/arthurkushman/go-hungarian"

	"github.com/vicdevcode/drivee-test/internal/entity"
)

func HungarianAlg(orders []entity.Order, couriers []entity.Courier) ([]entity.Order, []entity.Courier) {
	COMatrix := CreateCourierOrderMatrix(orders, couriers)
	deletedIds := []int{}

	// Цикл венгерского алгоритма
	for a := 0; a < len(orders); a += len(couriers) {
		sqMatrix := ConvertToSquareMatrix(COMatrix)

		// назначаем курьеров
		fmt.Println(sqMatrix)
		if len(sqMatrix) == 2 && sqMatrix[1][1] != 0 {
			for i := range couriers {
				id := i
				if len(deletedIds) >= len(couriers) {
					if deletedIds[0] <= id {
						id++
					}
					for i := 1; i < len(deletedIds)-len(deletedIds)%len(couriers); i++ {
						if deletedIds[i] <= id {
							id++
						}
					}
				}
				orders[id].CourierID = &couriers[i].ID
				couriers[couriers[i].ID].Orders = append(couriers[couriers[i].ID].Orders, orders[id])
			}
		}

		hung := hungarian.SolveMin(sqMatrix)
		fmt.Println(hung)
		// этот массив нужен для корректного удаления столбцов в COMatrix
		wannaBeDeletedIds := []int{}
		for courierId, path := range hung {
			if courierId >= len(couriers) {
				continue
			}
			for orderId, dist := range path {
				if dist <= 0.0 {
					continue
				}
				id := orderId
				// После удаления стобцов матрицы, заказы могут потерять свою
				// начальную очередность. Я проверяю, если индекс выше, чем первый
				// удаленный индекс, то он увеличивается.

				// Пример: 3x3 = [0, 1(удалили этот столбец), 2] => [0, 1]
				// Если индекс выше или равен удаленому индексу, то индекс увеличивается
				// один раз, тогда мы может определить изначальный порядок матрицы.
				// [0, X, 2]. Видим, что тут был удален 1, но 0 и 2 не были удалены и мы
				// можем корректно определить id заказов.
				if len(deletedIds) >= len(couriers) {
					if deletedIds[0] <= id {
						id++
					}
					for i := 1; i < len(deletedIds)-len(deletedIds)%len(couriers); i++ {
						if deletedIds[i] <= id {
							id++
						}
					}
				}
				deletedIds = append(deletedIds, id)
				wannaBeDeletedIds = append(wannaBeDeletedIds, orderId)
				orders[id].CourierID = &courierId
				couriers[courierId].Orders = append(couriers[courierId].Orders, orders[id])
			}
		}
		sort.Ints(wannaBeDeletedIds)
		for i := len(wannaBeDeletedIds) - 1; i >= 0; i-- {
			realId := wannaBeDeletedIds[i]
			COMatrix = RemoveVisitedOrder(COMatrix, realId)
		}
		sort.Ints(deletedIds)
	}

	return orders, couriers
}

func SetClusters(couriers []entity.Courier) []entity.Cluster {
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
