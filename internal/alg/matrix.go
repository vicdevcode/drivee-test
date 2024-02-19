package alg

import (
	"github.com/vicdevcode/drivee-test/internal/entity"
)

// Матрица, где курьеры - строки, а заказы - столбцы
func CreateCourierOrderMatrix(orders []entity.Order, couriers []entity.Courier) [][]float64 {
	matrix := [][]float64{}
	for _, courier := range couriers {
		subMatrix := []float64{}
		for _, order := range orders {
			subMatrix = append(subMatrix, calculateTime(courier.Speed, calculateDistance(courier.Coordinates, order.Start, order.End, courier.Penalty)))
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
		good := 0
		for j := range matrix[i] {
			if j != id {
				matrix[i][good] = matrix[i][j]
				good++
			}
		}
		matrix[i] = matrix[i][:good]
	}

	return matrix
}
