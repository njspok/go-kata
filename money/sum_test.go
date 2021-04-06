package money

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSum(t *testing.T) {
	t.Run("sum", func(t *testing.T) {
		sum := NewSum(NewDollar(3), NewDollar(4))
		bank := NewBank()
		result := bank.Reduce(sum, "USD")
		require.True(t, NewDollar(7).Equals(result))
	})
}
