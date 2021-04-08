package money

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSum(t *testing.T) {
	t.Run("reduce", func(t *testing.T) {
		sum := NewSum(NewDollar(3), NewDollar(4))
		bank := NewBank()
		result := bank.Reduce(sum, "USD")
		require.True(t, NewDollar(7).Equals(result))
	})
	t.Run("sum", func(t *testing.T) {
		bucks := NewDollar(1)
		frank := NewFrank(4)
		sum := NewSum(bucks, frank).Plus(bucks)
		bank := NewBank()
		bank.AddRate("CHF", "USD", 2)
		result := bank.Reduce(sum, "USD")
		require.True(t, NewDollar(4).Equals(result))
	})
	t.Run("times", func(t *testing.T) {
		bucks := NewDollar(1)
		frank := NewFrank(4)
		sum := NewSum(bucks, frank).Times(3)
		bank := NewBank()
		bank.AddRate("CHF", "USD", 2)
		result := bank.Reduce(sum, "USD")
		require.True(t, NewDollar(9).Equals(result))
	})
}
