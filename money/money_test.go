package money

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMoney(t *testing.T) {
	t.Run("compare different currency", func(t *testing.T) {
		require.False(t, NewFrank(5).Equals(NewDollar(5)))
		require.False(t, NewFrank(5).Equals(NewDollar(11)))

		require.False(t, NewDollar(5).Equals(NewFrank(5)))
		require.False(t, NewDollar(5).Equals(NewFrank(11)))
	})
}
