package genetic_alg

import (
	"fmt"
	"math/rand"
)

type SelectionType string

const (
	Roulette   SelectionType = "roulette"
	Tournament SelectionType = "tournament"
)

func New[T any](
	initPopulation []Chromosome,
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

// todo remove T type
type GeneticAlgorithm[T any] struct {
	population      []Chromosome
	threshold       float64
	maxGenerations  int
	mutationChance  float64
	crossoverChance float64
	selectionType   SelectionType
	fitnessKey      func() float64
}

func (ga *GeneticAlgorithm[T]) Run() Chromosome {
	best := maxFitness(ga.population)
	for generation := 0; generation < ga.maxGenerations; generation++ {
		if best.Fitness() >= ga.threshold {
			return best
		}

		fmt.Println("Generation", generation, "Best", best.Fitness(), "Avg", avg(ga.population)) // todo replace

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
	var newPopulation []Chromosome

	for len(newPopulation) < len(ga.population) {
		var parents [2]Chromosome
		if ga.selectionType == Roulette {
			// roulette
			parents = ga.pickRoulette(listFitness(ga.population))
		} else {
			// tournament
			parents = ga.pickTournament(len(ga.population) / 2)

		}

		if rand.Float64() < ga.crossoverChance {
			p1, p2 := parents[0].Crossover(parents[1])
			newPopulation = append(newPopulation, p1, p2)
		} else {
			newPopulation = append(newPopulation, parents[0], parents[1])
		}
	}

	if len(newPopulation) > len(ga.population) {
		newPopulation = newPopulation[:len(newPopulation)-1]
	}

	ga.population = newPopulation
}

func (ga *GeneticAlgorithm[T]) pickRoulette(list []float64) [2]Chromosome {
	return [2]Chromosome{}
}

func (ga *GeneticAlgorithm[T]) pickTournament(i int) [2]Chromosome {
	return [2]Chromosome{}
}

func listFitness(list []Chromosome) []float64 {
	var result []float64
	for _, individual := range list {
		result = append(result, individual.Fitness())
	}
	return result
}

func maxFitness(list []Chromosome) Chromosome {
	var result Chromosome
	for _, individual := range list {
		if result.Fitness() < individual.Fitness() {
			result = individual
		}
	}
	return result
}

func avg(list []Chromosome) float64 {
	var sum float64
	for _, individual := range list {
		sum += individual.Fitness()
	}
	return sum / float64(len(list))
}
