package lamp

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCounterWrapper(t *testing.T) {
	t.Run("create", func(t *testing.T) {
		wrapper := NewCounterWrapper(
			NewWorkLamp(),
		)
		require.NotNil(t, wrapper)
	})
	t.Run("counting", func(t *testing.T) {
		wrapper := NewCounterWrapper(
			NewWorkLamp(),
		)
		require.Zero(t, wrapper.Count())

		wrapper.Switch()
		require.Equal(t, uint(1), wrapper.Count())

		wrapper.Switch()
		require.Equal(t, uint(2), wrapper.Count())

		wrapper.Switch()
		require.Equal(t, uint(3), wrapper.Count())

		wrapper.Switch()
		require.Equal(t, uint(4), wrapper.Count())
	})
	t.Run("translating", func(t *testing.T) {
		lamp := NewWorkLamp()
		wrapper := NewCounterWrapper(lamp)

		wrapper.Switch()
		require.True(t, wrapper.IsLighted())
		require.True(t, lamp.IsLighted())

		wrapper.Switch()
		require.False(t, wrapper.IsLighted())
		require.False(t, lamp.IsLighted())

		wrapper.Switch()
		require.True(t, wrapper.IsLighted())
		require.True(t, lamp.IsLighted())
	})
}
