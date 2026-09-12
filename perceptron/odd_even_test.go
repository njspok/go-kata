package perceptron

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsOdd(t *testing.T) {
	for n := -100; n <= 100; n++ {
		require.Equal(t, n%2 != 0, IsOdd(n), n)
	}
}
