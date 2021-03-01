package lamp

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMan(t *testing.T) {
	t.Run("create", func(t *testing.T) {
		man := NewMan(1)
		require.NotNil(t, man)
	})
	t.Run("light on", func(t *testing.T) {
		man := NewMan(1)

		lamp := NewWorkLamp()
		require.False(t, lamp.IsLighted())

		// light on
		err := man.LightOn(lamp)
		require.NoError(t, err)
		require.True(t, lamp.IsLighted())

		// already lighted
		err = man.LightOn(lamp)
		require.NoError(t, err)
		require.True(t, lamp.IsLighted())
	})
	t.Run("already lighted", func(t *testing.T) {
		man := NewMan(1)

		// make lighted lamp
		lamp := NewWorkLamp()
		lamp.Switch()
		require.True(t, lamp.IsLighted())

		// already lighted
		err := man.LightOn(lamp)
		require.NoError(t, err)
		require.True(t, lamp.IsLighted())
	})
	t.Run("impatient", func(t *testing.T) {
		man := NewMan(0)

		lamp := NewCounterWrapper(
			NewBrokenLamp(10),
		)

		// only 1 attempt
		err := man.LightOn(lamp)
		require.ErrorIs(t, err, ErrLampNotWorking)
		require.Equal(t, uint(1), lamp.Count())
	})
	t.Run("patient", func(t *testing.T) {
		man := NewMan(5)

		lamp := NewCounterWrapper(
			NewBrokenLamp(10),
		)

		// trying
		err := man.LightOn(lamp)
		require.ErrorIs(t, err, ErrLampNotWorking)
		require.Equal(t, uint(6), lamp.Count())
	})
	t.Run("lucky", func(t *testing.T) {
		man := NewMan(5)

		lamp := NewCounterWrapper(
			NewBrokenLamp(3),
		)

		// trying
		err := man.LightOn(lamp)
		require.NoError(t, err)
		require.Equal(t, uint(3), lamp.Count())
	})
}
