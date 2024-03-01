package perceptron

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	t.Run("teacher", func(t *testing.T) {
		// login or element simulation
		sampleData := []SampleData{
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

		initWeights(p)

		for i := 0; i < 5; i++ {
			countMatched := 0
			for n, sample := range sampleData {
				matched, actual := lesson(p, sample.input, sample.expected)
				fmt.Printf("Test %d: actual %f expected %f\n", n, actual, sample.expected)
				if matched {
					countMatched++
				}
			}
			fmt.Printf("Matched %d from %d\n", countMatched, len(sampleData))
		}
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
