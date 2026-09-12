package perceptron

func IsOdd(n int) bool {
	d := lastDigit(n)
	input := digitToInput(d)
	return oddPerceptron.Evaluate(input) == 1
}

var oddPerceptron = newOddPerceptron()

func newOddPerceptron() *Perceptron {
	p := New()
	// вес 1 на нечётных цифрах 0–9
	p.SetWeights([]float64{0, 1, 0, 1, 0, 1, 0, 1, 0, 1})
	p.SetBiasWeight(0)
	return p
}

func lastDigit(n int) int {
	last := n % 10
	if last < 0 {
		last = -last
	}
	return last
}

func digitToInput(n int) []float64 {
	input := make([]float64, 10)
	input[n] = 1
	return input
}
