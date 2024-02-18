package alg

import (
	"math"

	"github.com/arthurkushman/go-hungarian"

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

// Матрица, где курьеры - строки, а заказы - столбцы
func CreateCourierOrderMatrix(orders []entity.Order, couriers []entity.Courier) [][]float64 {
	matrix := [][]float64{}
	for _, courier := range couriers {
		subMatrix := []float64{}
		for _, order := range orders {
			subMatrix = append(subMatrix, CalculateDistance(courier.Point, order.Start, order.End))
		}
		matrix = append(matrix, subMatrix)
	}
	return matrix
}

// Конвертируем матрицу в квадратичную матрицу, добавляя
// фиктивных курьеров
func ConvertToSquareMatrix(matrix [][]float64) [][]float64 {
	squareMatrix := [][]float64{}
	row := len(matrix)
	column := len(matrix[0])
	size := column
	if row > column {
		size = row
	}
	for i := 0; i < size; i++ {
		temp := []float64{}
		for j := 0; j < size; j++ {
			temp = append(temp, 0)
		}
		squareMatrix = append(squareMatrix, temp)
	}
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			squareMatrix[i][j] = matrix[i][j]
		}
	}

	return squareMatrix
}

// Убираем заказы, которые были назначены курьерам,
// заменяя все значения столбца заказа на максимальное число
func RemoveVisitedOrder(matrix [][]float64, id int) [][]float64 {
	for i := range matrix {
		for j := range matrix[i] {
			if j == id {
				matrix[i][j] = math.MaxFloat64
			}
		}
	}

	return matrix
}

func SetClusters() []entity.Cluster {
	orders, couriers := FakeData()

	// Создаем временный список, чтобы не изменять оригинал
	COMatrix := CreateCourierOrderMatrix(orders, couriers)
	_ = COMatrix

	tempCOMatrix := make([][]float64, len(COMatrix))
	copy(tempCOMatrix, COMatrix)

	for a := 0; a < len(orders); a += len(couriers) {
		sqMatrix := ConvertToSquareMatrix(tempCOMatrix)

		hung := hungarian.SolveMin(sqMatrix)
		for courierId, value := range hung {
			if courierId >= len(couriers) {
				continue
			}
			for orderId, dist := range value {
				if dist == 0 {
					continue
				}
				orders[orderId].CourierID = &courierId
				couriers[courierId].Orders = append(couriers[courierId].Orders, orders[orderId])
				tempCOMatrix = RemoveVisitedOrder(tempCOMatrix, orderId)
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
