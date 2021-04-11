package xunit

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWasRun(t *testing.T) {
	t.Run("exist", func(t *testing.T) {
		test := NewWasRun("TestMethod", t)
		require.False(t, test.wasRun)
		test.Run()
		require.True(t, test.wasRun)
	})
	t.Run("not exist", func(t *testing.T) {
		require.PanicsWithValue(
			t,
			"method not found",
			func() {
				test := NewWasRun("SomeOtherNotExistMethod", t)
				test.Run()
			},
		)
	})
}
