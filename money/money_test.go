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
	t.Run("plus", func(t *testing.T) {
		money := NewMoney(12, "USD")
		expr := money.Plus(money)
		bank := NewBank()
		reduced := bank.Reduce(expr, "USD")
		require.True(t, NewMoney(24, "USD").Equals(reduced))
	})
}

func TestDollar(t *testing.T) {
	t.Run("create", func(t *testing.T) {
		dollar := NewDollar(11)
		require.Equal(t, uint(11), dollar.Amount())
		require.Equal(t, "USD", dollar.Currency())
	})
}

func TestFrank(t *testing.T) {
	t.Run("create", func(t *testing.T) {
		frank := NewFrank(11)
		require.Equal(t, uint(11), frank.Amount())
		require.Equal(t, "CHF", frank.Currency())
	})
}
