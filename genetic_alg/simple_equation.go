package genetic_alg

import (
	"math"
	"math/rand"
)

func RandomSimpleEquation() *SimpleEquation {
	return &SimpleEquation{
		x: rand.Intn(100) - 50,
		y: rand.Intn(100) - 50,
	}
}

type SimpleEquation struct {
	x int
	y int
}

func (s *SimpleEquation) Fitness() float64 {
	z := s.x*s.x + s.y*s.y
	if z == 0 {
		return math.MaxFloat64
	}
	return 1 / float64(z)
}

func (s *SimpleEquation) Crossover(other Chromosome) (Chromosome, Chromosome) {
	child1 := s.copy()
	child2 := other.(*SimpleEquation).copy()
	child1.y = other.(*SimpleEquation).y
	child2.y = s.y
	return child1, child2
}

func (s *SimpleEquation) Mutate() {
	if rand.Float64() > 0.5 {
		if rand.Float64() > 0.5 {
			s.x += 1
		} else {
			s.x -= 1
		}
	} else {
		if rand.Float64() > 0.5 {
			s.y += 1
		} else {
			s.y -= 1
		}
	}
}

func (s *SimpleEquation) copy() *SimpleEquation {
	return &SimpleEquation{
		x: s.x,
		y: s.y,
	}
}
