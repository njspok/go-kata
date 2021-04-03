package money

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFrank(t *testing.T) {
	t.Run("amount", func(t *testing.T) {
		require.Equal(t, uint(11), NewFrank(11).Amount())
	})
	t.Run("equals", func(t *testing.T) {
		require.True(t, NewFrank(5).Equals(NewFrank(5)))
		require.False(t, NewFrank(5).Equals(NewFrank(11)))
	})
	t.Run("times", func(t *testing.T) {
		five := NewFrank(5)
		require.True(t, NewFrank(10).Equals(five.Times(2)))
		require.True(t, NewFrank(15).Equals(five.Times(3)))
	})
}
