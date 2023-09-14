package genetic_alg

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSimpleEquation(t *testing.T) {
	var initPop []Chromosome
	for i := 0; i < 20; i++ {
		initPop = append(initPop, RandomSimpleEquation())
	}

	ga := New(
		initPop,
		13,
		200,
		0.1,
		0.7,
		Roulette,
	)

	result := ga.Run()
	require.EqualValues(t, &SimpleEquation{
		x: 0,
		y: 0,
	}, result)
}
