package money

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDollar(t *testing.T) {
	t.Run("amount", func(t *testing.T) {
		require.Equal(t, uint(11), NewDollar(11).Amount())
	})
	t.Run("equals", func(t *testing.T) {
		require.True(t, NewDollar(5).Equals(NewDollar(5)))
		require.False(t, NewDollar(5).Equals(NewDollar(11)))
	})
	t.Run("times", func(t *testing.T) {
		five := NewDollar(5)
		require.True(t, NewDollar(10).Equals(five.Times(2)))
		require.True(t, NewDollar(15).Equals(five.Times(3)))
	})
}
