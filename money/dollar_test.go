package money

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDollar(t *testing.T) {
	t.Run("times", func(t *testing.T) {
		five := NewDollar(5)
		product := five.Times(2)
		require.Equal(t, uint(10), product.Amount())
		product = five.Times(3)
		require.Equal(t, uint(15), product.Amount())
	})
}
