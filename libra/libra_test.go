package libra

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLibra(t *testing.T) {
	t.Run("create", func(t *testing.T) {
		libra := NewLibra()
		require.NotNil(t, libra)
		require.Zero(t, libra.Balance())
		require.Zero(t, libra.Left())
		require.Zero(t, libra.Right())
	})
	t.Run("put", func(t *testing.T) {
		libra := NewLibra()
		libra.Put(1)
		require.Equal(t, uint(1), libra.Balance())
		require.Equal(t, uint(1), libra.Left())
		require.Equal(t, uint(0), libra.Right())
	})
	t.Run("balance", func(t *testing.T) {
		libra := NewLibra()
		libra.Put(1)
		libra.Put(1)
		require.Equal(t, uint(0), libra.Balance())
		require.Equal(t, uint(1), libra.Left())
		require.Equal(t, uint(1), libra.Right())
	})
	t.Run("left", func(t *testing.T) {
		libra := NewLibra()
		libra.Put(9)
		libra.Put(10)
		libra.Put(2)
		require.Equal(t, uint(1), libra.Balance())
		require.Equal(t, uint(11), libra.Left())
		require.Equal(t, uint(10), libra.Right())
	})
	t.Run("right", func(t *testing.T) {
		libra := NewLibra()
		libra.Put(9)
		libra.Put(10)
		libra.Put(2)
		libra.Put(3)
		require.Equal(t, uint(2), libra.Balance())
		require.Equal(t, uint(11), libra.Left())
		require.Equal(t, uint(13), libra.Right())
	})
	t.Run("left reduce", func(t *testing.T) {
		libra := NewLibra()
		libra.Put(1)
		libra.Put(10)
		libra.Put(1)
		libra.Put(1)
		require.Equal(t, uint(7), libra.Balance())
		require.Equal(t, uint(3), libra.Left())
		require.Equal(t, uint(10), libra.Right())
	})
	t.Run("right reduce", func(t *testing.T) {
		libra := NewLibra()
		libra.Put(10)
		libra.Put(1)
		libra.Put(1)
		libra.Put(1)
		require.Equal(t, uint(7), libra.Balance())
		require.Equal(t, uint(10), libra.Left())
		require.Equal(t, uint(3), libra.Right())
	})
}
