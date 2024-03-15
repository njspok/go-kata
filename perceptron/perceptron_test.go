package perceptron

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTraining(t *testing.T) {
	t.Run("success training OR logic", func(t *testing.T) {
		// logic OR element simulation
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

		err := Training(p, samples, 5)
		require.NoError(t, err)
	})
	t.Run("success training AND logic", func(t *testing.T) {
		// logic OR element simulation
		samples := []SampleData{
			{
				input:    []float64{0, 0},
				expected: 0,
			},
			{
				input:    []float64{0, 1},
				expected: 0,
			},
			{
				input:    []float64{1, 0},
				expected: 0,
			},
			{
				input:    []float64{1, 1},
				expected: 1,
			},
		}

		p := New()

		err := Training(p, samples, 10)
		require.NoError(t, err)
	})
	t.Run("fail traing OR logic", func(t *testing.T) {
		// logic OR element simulation
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

		err := Training(p, samples, 1)
		require.ErrorIs(t, err, ErrTrainingFailed)
	})
}

func TestPerceptron(t *testing.T) {
	p := New()

	p.SetWeights([]float64{1, 1, 0.5})
	p.SetBiasWeight(1)
	out := p.Evaluate([]float64{1, 1, 1})
	require.EqualValues(t, 1, out)

	p.SetWeights([]float64{0, 0, 0})
	p.SetBiasWeight(0.1)
	out = p.Evaluate([]float64{1, 1, 1})
	require.EqualValues(t, 0, out)
}
