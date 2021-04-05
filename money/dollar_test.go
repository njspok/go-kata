package money

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDollar(t *testing.T) {
	t.Run("create", func(t *testing.T) {
		dollar := NewDollar(11)
		require.Equal(t, uint(11), dollar.Amount())
		require.Equal(t, "USD", dollar.Currency())
	})
}
