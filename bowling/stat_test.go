package bowling

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewStat(t *testing.T) {
	t.Run("create", func(t *testing.T) {
		stat := NewStat()

		require.Len(t, stat.frames, 9)

		for i := 0; i < 9; i++ {
			require.Equal(t, Frame{}, stat.frames[i])
		}

		require.Equal(t, FinalFrame{}, stat.final)
	})
}
