package perceptron

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type SampleData struct {
	input  []float64
	result float64
}

func Test(t *testing.T) {
	t.Run("teacher", func(t *testing.T) {
		// login or element simulation
		data := []SampleData{
			{
				input:  []float64{0, 0},
				result: 0,
			},
			{
				input:  []float64{0, 1},
				result: 1,
			},
			{
				input:  []float64{1, 0},
				result: 1,
			},
			{
				input:  []float64{1, 1},
				result: 1,
			},
		}

		p := New()
		p.SetWeight([]float64{0, 0})
		p.SetBiasWeight(0)

		matched := 0
		for n, sample := range data {
			result := p.Run(sample.input)
			fmt.Printf("Test %d: actual %f expected %f\n", n, result, sample.result)
			if result == sample.result {
				matched++
			}
		}
		fmt.Printf("Matched %d from %d\n", matched, len(data))

	})
	t.Run("logic", func(t *testing.T) {
		p := New()

		p.SetWeight([]float64{1, 1, 0.5})
		p.SetBiasWeight(1)
		out := p.Run([]float64{1, 1, 1})
		require.EqualValues(t, 1, out)

		p.SetWeight([]float64{0, 0, 0})
		p.SetBiasWeight(0.1)
		out = p.Run([]float64{1, 1, 1})
		require.EqualValues(t, 0, out)
	})
}
