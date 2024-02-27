package perceptron

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	p := New()
	p.SetWeight([]float64{1, 1, 0.5})
	out := p.Run([]float64{1, 1, 1})
	require.EqualValues(t, 2.5, out)
}
