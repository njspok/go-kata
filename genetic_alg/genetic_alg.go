package genetic_alg

import (
	"fmt"
	"math"
	"math/rand"
)

type SelectionType string

const (
	Roulette   SelectionType = "roulette"
	Tournament SelectionType = "tournament"
)

func New[T Chromosome[T]](
	initPopulation []T,
	threshold float64,
	maxGenerations int,
	mutationChance float64,
	crossoverChance float64,
	selectionType SelectionType,
) *GeneticAlgorithm[T] {
	return &GeneticAlgorithm[T]{
		population:      initPopulation,
		threshold:       threshold,
		maxGenerations:  maxGenerations,
		mutationChance:  mutationChance,
		crossoverChance: crossoverChance,
		selectionType:   selectionType,
		fitnessKey:      initPopulation[0].Fitness,
	}
}

type GeneticAlgorithm[T Chromosome[T]] struct {
	population      []T
	threshold       float64
	maxGenerations  int
	mutationChance  float64
	crossoverChance float64
	selectionType   SelectionType
	fitnessKey      func() float64
}

func (ga *GeneticAlgorithm[T]) Run() T {
	best := maxFitness(ga.population)
	for generation := 0; generation < ga.maxGenerations; generation++ {
		if best.Fitness() >= ga.threshold {
			return best
		}

		fmt.Println(
			"Generation",
			generation,
			"Best",
			best.Fitness(),
			"Avg",
			avg(ga.population),
		)

		ga.reproduceAndReplace()
		ga.mutate()

		highest := maxFitness(ga.population)
		if highest.Fitness() > best.Fitness() {
			best = highest
		}
	}
	return best
}

func (ga *GeneticAlgorithm[T]) mutate() {
	for _, individual := range ga.population {
		if rand.Float64() < ga.mutationChance {
			individual.Mutate()
		}
	}
}

func (ga *GeneticAlgorithm[T]) reproduceAndReplace() {
	var newPopulation []T

	for len(newPopulation) < len(ga.population) {
		var p1 T
		var p2 T

		if ga.selectionType == Roulette {
			// roulette
			p1, p2 = ga.pickRoulette(listFitness(ga.population))
		} else {
			// tournament
			p1, p2 = ga.pickTournament(len(ga.population) / 2)

		}

		if rand.Float64() < ga.crossoverChance {
			p1, p2 := p1.Crossover(p2)
			newPopulation = append(newPopulation, p1, p2)
		} else {
			newPopulation = append(newPopulation, p1, p2)
		}
	}

	if len(newPopulation) > len(ga.population) {
		newPopulation = newPopulation[:len(newPopulation)-1]
	}

	ga.population = newPopulation
}

func (ga *GeneticAlgorithm[T]) pickRoulette(list []float64) (T, T) {
	r := choices(ga.population, list, 2)
	return r[0], r[1]
}

func (ga *GeneticAlgorithm[T]) pickTournament(i int) (T, T) {
	// participants: List[C] = choices(self._population, k=num_participants)
	// return tuple(nlargest(2, participants, key=self._fitness_key))
	panic("need implement")
}

func listFitness[T Chromosome[T]](list []T) []float64 {
	var result []float64
	for _, individual := range list {
		result = append(result, individual.Fitness())
	}
	return result
}

func maxFitness[T Chromosome[T]](list []T) T {
	result := list[0]
	for _, individual := range list {
		if result.Fitness() < individual.Fitness() {
			result = individual
		}
	}
	return result
}

func avg[T Chromosome[T]](list []T) float64 {
	var sum float64
	for _, individual := range list {
		sum += individual.Fitness()
	}
	return sum / float64(len(list))
}

func choices[T any](list []T, weights []float64, k int) []T {
	var sum float64
	for _, w := range weights {
		sum += math.Abs(w)
	}

	result := make([]T, k)
	for i := 0; i < k; i++ {
		r := rand.Float64() * sum
		for j, w := range weights {
			r -= math.Abs(w)
			if r <= 0 {
				result[i] = list[j]
				break
			}
		}
	}
	return result
}
