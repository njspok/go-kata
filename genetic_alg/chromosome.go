package genetic_alg

type Chromosome interface {
	Fitness() float64
	RandomInstance(cls Chromosome) Chromosome
	Crossover(other Chromosome) (Chromosome, Chromosome)
	Mutate()
}
