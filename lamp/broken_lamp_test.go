package lamp

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBrokenLamp(t *testing.T) {
	t.Run("create", func(t *testing.T) {
		lamp := NewBrokenLamp(10)
		require.NotNil(t, lamp)
	})
	t.Run("no lighting", func(t *testing.T) {
		lamp := NewBrokenLamp(10)
		require.False(t, lamp.IsLighted())

		// no
		lamp.Switch()
		require.False(t, lamp.IsLighted())

		// no
		lamp.Switch()
		require.False(t, lamp.IsLighted())
	})
	t.Run("lighting", func(t *testing.T) {
		lamp := NewBrokenLamp(3)
		require.False(t, lamp.IsLighted())

		// success light cycle
		for i := 0; i < 5; i++ {
			// try 1
			lamp.Switch()
			require.False(t, lamp.IsLighted())

			// try 2
			lamp.Switch()
			require.False(t, lamp.IsLighted())

			// try 3
			lamp.Switch()
			require.True(t, lamp.IsLighted())
		}
	})
}
