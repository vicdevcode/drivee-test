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

func CreateGenome(V int) []int {
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

	return genome
}

func MutatedGen(genome []int, V int) []int {
	mutGenome := make([]int, len(genome))
	copy(mutGenome, genome)
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
	return mutGenome
}

func CalFitness(genome []int, mp [][]float64) float64 {
	f := 0.0
	for i := 0; i < len(genome)-1; i++ {
		if mp[genome[i]][genome[i+1]] == MAX {
			return MAX
		}
		f += mp[genome[i]][genome[i+1]]
	}
	return f
}

func cooldown(temp float64) float64 {
	return (90 * temp) / 100
}

func SolveGA(V int, popSize int, mp [][]float64) (float64, []int) {
	gen := 1

	genThres := 100
	population := []individual{}
	temp := individual{
		genome:  []int{},
		fitness: 0,
	}

	for i := 0; i < popSize; i++ {
		temp.genome = CreateGenome(V)
		temp.fitness = CalFitness(temp.genome, mp)
		population = append(population, temp)
	}

	temperature := 10000.0

	for temperature > 1000 && gen <= genThres {
		sort.Sort(ByFitness(population))
		newPopulation := []individual{}

		for i := 0; i < popSize; i++ {
			p1 := population[i]

			for {
				newG := MutatedGen(p1.genome, V)
				newGenome := individual{}
				newGenome.genome = newG
				newGenome.fitness = CalFitness(newGenome.genome, mp)

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
