package alg

import (
	"math"

	"github.com/vicdevcode/drivee-test/internal/entity"
)

// Измеряем расстояние между 3 точками (C -> A -> B) либо (B -> A -> B).
// Почему B в начале? Потому что B является концом заказа,
// где окажется в итоге курьер
func CalculateDistance(First entity.Point, Start entity.Point, End entity.Point) float64 {
	fromFirstToStart := math.Sqrt(math.Pow(First.X-Start.X, 2) + math.Pow(First.Y-Start.Y, 2))
	fromStartToEnd := math.Sqrt(math.Pow(End.X-Start.X, 2) + math.Pow(End.Y-Start.Y, 2))
	return fromFirstToStart + fromStartToEnd
}

func FakeData() ([]entity.Order, []entity.Courier) {
	return []entity.Order{
			{
				Start: entity.Point{X: 3, Y: 18},
				End:   entity.Point{X: 15, Y: 8},
				Price: 50,
			},
			{
				Start: entity.Point{X: 9, Y: 4},
				End:   entity.Point{X: 13, Y: 5},
				Price: 50,
			},
			{
				Start: entity.Point{X: 6, Y: 8},
				End:   entity.Point{X: 14, Y: 13},
				Price: 50,
			},
			{
				Start: entity.Point{X: 4, Y: 13},
				End:   entity.Point{X: 8, Y: 11},
				Price: 50,
			},
			{
				Start: entity.Point{X: 3, Y: 2},
				End:   entity.Point{X: 16, Y: 2},
				Price: 50,
			},
			{
				Start: entity.Point{X: 10, Y: 18},
				End:   entity.Point{X: 11, Y: 12},
				Price: 50,
			},
			{
				Start: entity.Point{X: 2, Y: 4},
				End:   entity.Point{X: 17, Y: 10},
				Price: 50,
			},
		}, []entity.Courier{
			{
				ID:     0,
				Point:  entity.Point{X: 2, Y: 8},
				Orders: nil,
			},
			{
				ID:     1,
				Point:  entity.Point{X: 13, Y: 16},
				Orders: nil,
			},
		}
}

func SetClusters() []entity.Cluster {
	orders, couriers := FakeData()

	// Создаем временный список, чтобы не изменять оригинал
	tempOrders := make([]entity.Order, len(orders))
	copy(tempOrders, orders)

	// Алгоритм распределения
	// Цикл while. Пока не удаляться все элементы tempOrders
	// Будем дальше их распределять для каждого курьера
	for len(tempOrders) > 0 {
		for i := 0; i < len(couriers); i++ {

			// кандидат на минимальную дистанцию между курьером и заказом
			minCandidate := tempOrders[0]
			minCandidateDistance := CalculateDistance(
				couriers[i].Point,
				tempOrders[0].Start,
				tempOrders[0].End,
			)

			if len(tempOrders) == 1 {
				tempOrders[0].CourierID = &couriers[i].ID
				couriers[i].Orders = append(couriers[i].Orders, tempOrders[0])
				tempOrders = nil
				break
			}

			index := 0

			for j, order := range tempOrders {
				candidateDist := CalculateDistance(couriers[i].Point, order.Start, order.End)
				if minCandidateDistance > candidateDist {
					minCandidate = order
					minCandidateDistance = candidateDist
					index = j
				}
			}

			minCandidate.CourierID = &couriers[i].ID
			couriers[i].Orders = append(couriers[i].Orders, minCandidate)
			if len(tempOrders) > 1 {
				tempOrders[index] = tempOrders[len(tempOrders)-1]
				tempOrders[len(tempOrders)-1] = entity.Order{}
				tempOrders = tempOrders[:len(tempOrders)-1]
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
			courierRow = append(courierRow, CalculateDistance(courier.Point, order.Start, order.End))
		}
		cluster = append(cluster, courierRow)

		// Заполняем двумерную матрицу
		for i, order := range courier.Orders {
			cityRow := []float64{CalculateDistance(courier.Point, order.Start, order.End)}
			for j, orderToOrder := range courier.Orders {
				if j == i {
					cityRow = append(cityRow, 0)
					continue
				}
				cityRow = append(cityRow, CalculateDistance(
					order.End,
					orderToOrder.Start,
					orderToOrder.End,
				))
			}
			cluster = append(cluster, cityRow)
		}
		clusters = append(clusters, entity.Cluster{CourierID: courier.ID, Cluster: cluster})
	}

	return clusters
}
