package perceptron

func New() *Perceptron {
	return &Perceptron{
		weights:   []float64{},
		threshold: 0.5,
	}
}

func (p *Perceptron) Run(input []float64) float64 {
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

func (p *Perceptron) SetWeight(weights []float64) {
	p.weights = weights
}

func (p *Perceptron) SetBiasWeight(w float64) {
	p.biasWeight = w
}

type Perceptron struct {
	weights    []float64
	biasWeight float64
	threshold  float64
}
