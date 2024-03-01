package perceptron

import (
	"fmt"
	"math/rand"
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

		weights := []float64{
			rand.Float64() / 1000,
			rand.Float64() / 1000,
		}
		biasWeight := rand.Float64() / 1000

		fmt.Printf("weights: %v\n", weights)
		fmt.Printf("bias weight: %v\n", biasWeight)

		p.SetWeights(weights)
		p.SetBiasWeight(biasWeight)

		for i := 0; i < 5; i++ {
			matched := 0
			for n, sample := range sampleData {
				actual := p.Evaluate(sample.input)
				fmt.Printf("Test %d: actual %f expected %f\n", n, actual, sample.expected)
				if actual == sample.expected {
					matched++
				}

				// new weights
				var newWeights []float64
				for j := 0; j < p.InputCounts(); j++ {
					w := p.Weight(j) + (sample.expected-actual)*sample.input[j]
					newWeights = append(newWeights, w)
				}
				p.SetWeights(newWeights)

				newBiasWeight := p.BiasWeight() + (sample.expected-actual)*1
				p.SetBiasWeight(newBiasWeight)
			}
			fmt.Printf("Matched %d from %d\n", matched, len(sampleData))
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
