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
	t.Run("immutable", func(t *testing.T) {
		// Arrange
		cart := NewCart().Add("book", "123456")

		item, present := cart.Find("123456").Get()
		require.True(t, present)

		// Act
		item.qty = 99

		// Assert
		itemInCart, stillPresent := cart.Find("123456").Get()
		require.True(t, stillPresent)

		require.EqualValues(t, 1, itemInCart.Qty())
	})
}

func TestCopy(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		require.Nil(t, Copy[*Item](nil))
	})
	t.Run("value type", func(t *testing.T) {
		// Arrange
		item := Item{
			name: "book",
			sid:  "123456",
			qty:  1,
		}

		// Act
		newItem := Copy(item)

		// Assert
		require.Equal(t, item.Name(), newItem.Name())
		require.Equal(t, item.Qty(), newItem.Qty())
		require.Equal(t, item.Sid(), newItem.Sid())

		item.IncQty()

		require.NotEqual(t, item.Qty(), newItem.Qty())
	})
	t.Run("pointer type", func(t *testing.T) {
		// Arrange
		item := &Item{
			name: "book",
			sid:  "123456",
			qty:  1,
		}

		// Act
		newItem := Copy(item)

		// Assert
		require.Equal(t, item.Name(), newItem.Name())
		require.Equal(t, item.Qty(), newItem.Qty())
		require.Equal(t, item.Sid(), newItem.Sid())

		item.IncQty()

		require.NotEqual(t, item.Qty(), newItem.Qty())
	})
}
