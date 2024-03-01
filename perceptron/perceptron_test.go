package perceptron

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	t.Run("teacher", func(t *testing.T) {
		// login or element simulation
		samples := []SampleData{
			{
				input:    []float64{0, 0},
				expected: 0,
			},
			{
				input:    []float64{0, 1},
				expected: 1,
			},
			{
				input:    []float64{1, 0},
				expected: 1,
			},
			{
				input:    []float64{1, 1},
				expected: 1,
			},
		}

		p := New()

		train(p, samples)
	})
	t.Run("logic", func(t *testing.T) {
		p := New()

		p.SetWeights([]float64{1, 1, 0.5})
		p.SetBiasWeight(1)
		out := p.Evaluate([]float64{1, 1, 1})
		require.EqualValues(t, 1, out)

		p.SetWeights([]float64{0, 0, 0})
		p.SetBiasWeight(0.1)
		out = p.Evaluate([]float64{1, 1, 1})
		require.EqualValues(t, 0, out)
	})
}
