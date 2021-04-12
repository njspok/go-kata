package xunit

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWasRun(t *testing.T) {
	t.Run("exist", func(t *testing.T) {
		test := NewWasRun("TestMethod", t)
		require.Zero(t, test.log)
		test.Run()
		require.Equal(t, []string{"SetUp", "TestMethod"}, test.log)
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
