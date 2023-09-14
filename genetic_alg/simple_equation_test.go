package genetic_alg

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSimpleEquation(t *testing.T) {
	var initPop []*SimpleEquation
	for i := 0; i < 20; i++ {
		initPop = append(initPop, RandomSimpleEquation())
	}

	ga := New[*SimpleEquation](
		initPop,
		13,
		200,
		0.1,
		0.7,
		Roulette,
	)

	result := ga.Run()
	require.Equal(t, &SimpleEquation{
		x: 0,
		y: 0,
	}, result)
}
