package money

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDollar(t *testing.T) {
	t.Run("create", func(t *testing.T) {
		require.Equal(t, uint(11), NewDollar(11).Amount())
	})
	t.Run("times", func(t *testing.T) {
		five := NewDollar(5)
		product := five.Times(2)
		require.Equal(t, uint(10), product.Amount())
		product = five.Times(3)
		require.Equal(t, uint(15), product.Amount())
	})
	t.Run("equals", func(t *testing.T) {
		require.True(t, NewDollar(5).Equals(NewDollar(5)))
		require.False(t, NewDollar(5).Equals(NewDollar(11)))
	})
}
