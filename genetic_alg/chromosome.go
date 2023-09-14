package genetic_alg

type Chromosome[T any] interface {
	Fitness() float64
	Crossover(other T) (T, T)
	Mutate()
}
