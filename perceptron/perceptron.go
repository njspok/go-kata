package perceptron

import (
	"math/rand"
	"slices"
)

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

type SampleData struct {
	input    []float64
	expected float64
}

func initWeights(p *Perceptron) {
	weights := []float64{
		rand.Float64() / 1000,
		rand.Float64() / 1000,
	}
	biasWeight := rand.Float64() / 1000

	p.SetWeights(weights)
	p.SetBiasWeight(biasWeight)
}

func lesson(p *Perceptron, input []float64, expected float64) (matched bool, actual float64) {
	actual = p.Evaluate(input)
	if actual == expected {
		return true, actual
	}

	// new weights
	var newWeights []float64
	for j := 0; j < p.InputCounts(); j++ {
		w := p.Weight(j) + (expected-actual)*input[j]
		newWeights = append(newWeights, w)
	}
	p.SetWeights(newWeights)

	newBiasWeight := p.BiasWeight() + (expected-actual)*1
	p.SetBiasWeight(newBiasWeight)

	return false, actual
}
