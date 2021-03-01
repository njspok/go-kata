package lamp

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWorkLamp(t *testing.T) {
	t.Run("create", func(t *testing.T) {
		lamp := NewWorkLamp()
		require.NotNil(t, lamp)
	})
	t.Run("is lighted off", func(t *testing.T) {
		lamp := NewWorkLamp()
		require.False(t, lamp.IsLighted())
	})
	t.Run("switch", func(t *testing.T) {
		lamp := NewWorkLamp()
		require.False(t, lamp.IsLighted())

		// turn on
		lamp.Switch()
		require.True(t, lamp.IsLighted())

		// turn off
		lamp.Switch()
		require.False(t, lamp.IsLighted())

		// turn on
		lamp.Switch()
		require.True(t, lamp.IsLighted())
	})
}
