package bowling

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewFrames(t *testing.T) {
	t.Run("create", func(t *testing.T) {
		frames := NewFrames()

		require.Len(t, frames.frames, 10)

		for i := uint(1); i <= 10; i++ {
			require.Equal(t, i, frames.Frame(i).Number())
		}
	})
}
