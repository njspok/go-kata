package money

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFrank(t *testing.T) {
	t.Run("create", func(t *testing.T) {
		frank := NewFrank(11)
		require.Equal(t, uint(11), frank.Amount())
		require.Equal(t, "CHF", frank.Currency())
	})
}
