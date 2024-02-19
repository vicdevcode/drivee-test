package alg

import (
	"math"
	"math/rand"
	"sort"
)

const (
	MAX = math.MaxFloat64
)

type individual struct {
	genome  []int
	fitness float64
}

type ByFitness []individual

func (a ByFitness) Len() int           { return len(a) }
func (a ByFitness) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByFitness) Less(i, j int) bool { return a[i].fitness < a[j].fitness }

func randNum(start, end int) int {
	return rand.Intn(end-start) + start
}

func repeat(genome []int, temp int) bool {
	for _, gen := range genome {
		if gen == temp {
			return true
		}
	}
	return false
}

func (individ *individual) createGenome(V int) {
	genome := []int{0}
	for {
		if len(genome) == V {
			break
		}

		temp := randNum(1, V)
		if !repeat(genome, temp) {
			genome = append(genome, temp)
		}
	}
	individ.genome = genome
}

func (individ *individual) mutatedGen(V int) {
	mutGenome := make([]int, len(individ.genome))
	copy(mutGenome, individ.genome)
	for {
		r := randNum(1, V)
		r1 := randNum(1, V)
		if r1 != r {
			temp := mutGenome[r]
			mutGenome[r] = mutGenome[r1]
			mutGenome[r1] = temp
			break
		}
	}
	individ.genome = mutGenome
}

func (individ *individual) calFitness(mp [][]float64) {
	f := 0.0
	for i := 0; i < len(individ.genome)-1; i++ {
		if mp[individ.genome[i]][individ.genome[i+1]] == MAX {
			individ.fitness = MAX
			return
		}
		f += mp[individ.genome[i]][individ.genome[i+1]]
	}
	individ.fitness = f
}

func cooldown(temp float64) float64 {
	return (90 * temp) / 100
}

func SolveGA(V int, popSize int, mp [][]float64) (float64, []int) {
	gen := 1

	genThres := 1000
	population := []individual{}
	temp := individual{
		genome:  []int{},
		fitness: 0,
	}

	for i := 0; i < popSize; i++ {
		temp.createGenome(V)
		temp.calFitness(mp)
		population = append(population, temp)
	}

	temperature := 100000000.0

	for temperature > 100 && gen <= genThres {
		sort.Sort(ByFitness(population))
		newPopulation := []individual{}

		for i := 0; i < popSize; i++ {
			p1 := population[i]

			for {
				newGenome := individual{}
				newGenome.genome = p1.genome
				newGenome.mutatedGen(V)
				newGenome.calFitness(mp)

				if newGenome.fitness <= population[i].fitness {
					newPopulation = append(newPopulation, newGenome)
					break
				} else {
					prob := math.Pow(2.7, -1*((newGenome.fitness-population[i].fitness)/temperature))

					if prob > 0.5 {
						newPopulation = append(newPopulation, newGenome)
						break
					}
				}
			}
		}
		temperature = cooldown(temperature)
		population = newPopulation

		gen += 1
	}

	sort.Sort(ByFitness(population))

	return population[0].fitness, population[0].genome
}
