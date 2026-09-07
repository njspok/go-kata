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

func TestItem_SetQty(t *testing.T) {
	item := NewItem("book", "123456")

	newItem := item.SetQty(99)

	require.EqualValues(t, uint(1), item.Qty())
	require.EqualValues(t, 99, newItem.Qty())
}
