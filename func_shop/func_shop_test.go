package func_shop

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewItem(t *testing.T) {
	item := NewItem("book", "123456", 0)
	require.Equal(t, "book", item.Name())
	require.EqualValues(t, "123456", item.Sid())
	require.EqualValues(t, 1, item.qty)
}

func TestCart_Find(t *testing.T) {
	t.Run("already append", func(t *testing.T) {
		// Arrange
		cart := NewCart().Add("book", "123456", 11)

		// Act
		cart = cart.Add("book", "123456", 11)

		// Assert
		itemInCart, present := cart.Find("123456").Get()
		require.True(t, present)
		require.EqualValues(t, 2, itemInCart.Qty())
		require.Equal(t, 1, cart.Count())
	})
}

func TestCart_Add(t *testing.T) {
	t.Run("new item", func(t *testing.T) {
		// Arrange
		cart := NewCart()

		// Act
		newCart := cart.Add("book", "123456", 11)

		// Assert
		require.Equal(t, 0, cart.Count())
		require.True(t, cart.Find("123456").IsAbsent())

		require.Equal(t, 1, newCart.Count())
		item, present := newCart.Find("123456").Get()
		require.True(t, present)
		require.Equal(t, "book", item.Name())
		require.EqualValues(t, "123456", item.Sid())
		require.EqualValues(t, 1, item.Qty())
	})
	t.Run("change item price", func(t *testing.T) {
		// Arrange
		cart := NewCart().Add("book", "123456", 11)

		// Act
		cart = cart.Add("book", "123456", 22)

		// Assert
		require.Equal(t, 1, cart.Count())

		item, present := cart.Find("123456").Get()
		require.True(t, present)
		require.EqualValues(t, 22, item.Price())
		require.EqualValues(t, 2, item.Qty())
	})
	t.Run("existing item", func(t *testing.T) {
		// Arrange
		cart := NewCart().Add("book", "123456", 11)

		// Act
		newCart := cart.Add("book", "123456", 11)

		// Assert
		item, present := cart.Find("123456").Get()
		require.True(t, present)
		require.EqualValues(t, 1, item.Qty())
		require.Equal(t, 1, cart.Count())

		newItem, present := newCart.Find("123456").Get()
		require.True(t, present)
		require.EqualValues(t, 2, newItem.Qty())
		require.Equal(t, 1, newCart.Count())
	})
}

func TestCart_Del(t *testing.T) {
	t.Run("exist item", func(t *testing.T) {
		// Arrange
		cart := NewCart()
		cart = cart.Add("book", "123456", 11)

		// Act
		newCart := cart.Del("123456")

		// Assert
		require.Equal(t, 1, cart.Count())
		require.True(t, cart.Find("123456").IsPresent())

		require.Equal(t, 0, newCart.Count())
		require.True(t, newCart.Find("123456").IsAbsent())
	})
	t.Run("not exist item", func(t *testing.T) {
		// Arrange
		cart := NewCart()
		cart = cart.Add("book", "123456", 11)

		// Act
		newCart := cart.Del("666")

		// Assert
		require.Equal(t, 1, cart.Count())
		require.True(t, cart.Find("123456").IsPresent())

		require.Equal(t, 1, newCart.Count())
		require.True(t, newCart.Find("123456").IsPresent())
	})
	t.Run("empty cart", func(t *testing.T) {
		// Arrange
		cart := NewCart()

		// Act
		newCart := cart.Del("666")

		// Assert
		require.Equal(t, 0, cart.Count())
		require.Equal(t, 0, newCart.Count())
	})
}

func TestItem_SetPrice(t *testing.T) {
	t.Run("immutable", func(t *testing.T) {
		// Arrange
		item := NewItem("book", "123456", 11)

		// Act
		newItem := item.SetPrice(22)

		// Assert
		require.EqualValues(t, 22, newItem.Price())
		require.EqualValues(t, 11, item.Price())
	})
}
