package genetic_alg

type Chromosome interface {
	Fitness() float64
	Crossover(other Chromosome) (Chromosome, Chromosome)
	Mutate()
}
