package fractional_calc

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseExpression(t *testing.T) {
	t.Run("simple numbers", func(t *testing.T) {
		a, b, c, op, err := ParseExpression(`2+3=5`)
		require.NoError(t, err)
		require.Equal(t, Number(2), a)
		require.Equal(t, Number(3), b)
		require.Equal(t, Number(5), c)
		require.Equal(t, "+", op)
	})
	t.Run("fractions", func(t *testing.T) {
		a, b, c, op, err := ParseExpression(`2|3-4|5=-2|15`)
		require.NoError(t, err)
		require.Equal(t, Frac(2, 3), a)
		require.Equal(t, Frac(4, 5), b)
		require.Equal(t, Frac(-2, 15), c)
		require.Equal(t, "-", op)
	})
	t.Run("invalid string", func(t *testing.T) {
		_, _, _, _, err := ParseExpression("invalid string")
		require.Error(t, err)
	})
}
