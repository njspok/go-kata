package perceptron

func New() *Perceptron {
	return &Perceptron{
		weights: []float64{},
	}
}

func (p *Perceptron) Run(input []float64) float64 {
	var out float64
	for n, i := range input {
		out += i * p.weights[n]
	}
	return out
}

func (p *Perceptron) SetWeight(weights []float64) {
	p.weights = weights
}

type Perceptron struct {
	weights []float64
}
