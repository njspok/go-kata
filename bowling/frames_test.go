package bowling

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewFrames(t *testing.T) {
	t.Run("create", func(t *testing.T) {
		frames := NewFrames()
		require.Zero(t, frames.Len())
	})
	// todo need tests
}
