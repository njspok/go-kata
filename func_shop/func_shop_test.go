package func_shop

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewItem(t *testing.T) {
	item := NewItem("book", "123456")
	require.Equal(t, "book", item.Name())
	require.EqualValues(t, "123456", item.Sid())
	require.EqualValues(t, 1, item.qty)
}

func TestCart_Find(t *testing.T) {
	t.Run("already append", func(t *testing.T) {
		// Arrange
		cart := NewCart().Add("book", "123456")

		// Act
		cart.Add("book", "123456")

		// Assert
		itemInCart, present := cart.Find("123456").Get()
		require.True(t, present)
		require.EqualValues(t, 2, itemInCart.Qty())
		require.Equal(t, 1, cart.Count())
	})
}
