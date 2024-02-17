package alg

import "fmt"

type Distances [][]float64

// Объяснение подробное этой функции в README.md
func GenFakeDists(dists [][]float64, exit int) [][]float64 {
	n := len(dists)
	bigN := 10000000.0

	// Создаем матрицу с фиктивной точкой
	fakeDists := make([][]float64, n+1)
	for i := range fakeDists {
		fakeDists[i] = make([]float64, n+1)
	}

	// Определяем входную точку, а также
	// остальные поля заполняем "стеной"
	// в виде огромного числа
	for i := 0; i < n+1; i++ {
		if i == 1 {
			fakeDists[i][0] = 0
			continue
		}
		fakeDists[i][0] = bigN
	}

	// Определяем выходную точку
	for j := 0; j < n+1; j++ {
		if j == exit+1 {
			fakeDists[j][0] = 0
			continue
		}
		fakeDists[0][j] = bigN
	}

	// Копируем оригинал во внутрь
	// матрицы с фиктивной точкой
	for i := 1; i < n+1; i++ {
		for j := 1; j < n+1; j++ {
			fakeDists[i][j] = dists[i-1][j-1]
		}
	}

	return fakeDists
}

func HeldKapr(dists Distances) (float64, []int) {
	fmt.Println(dists)
	fakeDists := GenFakeDists(dists, 2)
	fmt.Println(fakeDists)
	return 0, []int{}
}
