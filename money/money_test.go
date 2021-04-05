package money

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMoney(t *testing.T) {
	t.Run("create", func(t *testing.T) {
		money := NewMoney(112, "USD")
		require.Equal(t, uint(112), money.Amount())
		require.Equal(t, "USD", money.Currency())
	})
	t.Run("equals different currency", func(t *testing.T) {
		require.False(t, NewMoney(5, "USD").Equals(NewMoney(5, "CHF")))
		require.False(t, NewMoney(5, "USD").Equals(NewMoney(11, "CHF")))
	})
	t.Run("equals different amount", func(t *testing.T) {
		require.True(t, NewMoney(5, "USD").Equals(NewMoney(5, "USD")))
		require.False(t, NewMoney(5, "USD").Equals(NewMoney(11, "USD")))
	})
	t.Run("times", func(t *testing.T) {
		five := NewMoney(5, "USD")
		require.True(t, NewMoney(10, "USD").Equals(five.Times(2)))
		require.True(t, NewMoney(15, "USD").Equals(five.Times(3)))
	})
}
