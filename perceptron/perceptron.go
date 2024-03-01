package perceptron

import "slices"

const defaultThreshold = 0.5

func New() *Perceptron {
	return &Perceptron{
		weights:   []float64{},
		threshold: defaultThreshold,
	}
}

func (p *Perceptron) Evaluate(input []float64) float64 {
	var out float64
	for n, i := range input {
		out += i * p.weights[n]
	}

	out += p.biasWeight * 1

	if out > p.threshold {
		return 1
	}

	return 0
}

func (p *Perceptron) InputCounts() int {
	return len(p.weights)
}

func (p *Perceptron) Weight(n int) float64 {
	return p.weights[n]
}

func (p *Perceptron) BiasWeight() float64 {
	return p.biasWeight
}

func (p *Perceptron) SetWeights(weights []float64) {
	p.weights = slices.Clone(weights)
}

func (p *Perceptron) SetWeight(n int, w float64) {
	p.weights[n] = w
}

func (p *Perceptron) SetBiasWeight(w float64) {
	p.biasWeight = w
}

type Perceptron struct {
	weights    []float64
	biasWeight float64
	threshold  float64
}
