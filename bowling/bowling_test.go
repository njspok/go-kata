package bowling

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBowling(t *testing.T) {
	t.Run("create", func(t *testing.T) {
		game := NewBowling()
		require.NotNil(t, game)
	})
	t.Run("roll", func(t *testing.T) {
		t.Run("all frames opened", func(t *testing.T) {
			game := NewBowling()
			require.Zero(t, game.Total())

			// 1
			game.Roll(1)
			game.Roll(2)
			require.Equal(t, uint(3), game.Total())

			// 2
			game.Roll(3)
			game.Roll(2)
			require.Equal(t, uint(8), game.Total())

			// 3
			game.Roll(3)
			game.Roll(4)
			require.Equal(t, uint(15), game.Total())

			// 4
			game.Roll(5)
			game.Roll(1)
			require.Equal(t, uint(21), game.Total())

			// 5
			game.Roll(3)
			game.Roll(4)
			require.Equal(t, uint(28), game.Total())

			// 6
			game.Roll(5)
			game.Roll(0)
			require.Equal(t, uint(33), game.Total())

			// 7
			game.Roll(2)
			game.Roll(3)
			require.Equal(t, uint(38), game.Total())

			// 8
			game.Roll(4)
			game.Roll(3)
			require.Equal(t, uint(45), game.Total())

			// 9
			game.Roll(2)
			game.Roll(2)
			require.Equal(t, uint(49), game.Total())

			// 10
			game.Roll(7)
			game.Roll(2)
			err := game.Roll(5)
			require.ErrorIs(t, err, ErrGameFinished)
			require.Equal(t, uint(58), game.Total())
			require.True(t, game.IsFinished())
		})
		t.Run("all strikes", func(t *testing.T) {
			// todo
		})
		t.Run("all spares", func(t *testing.T) {
			// todo
		})
		t.Run("mixed", func(t *testing.T) {
			// todo
		})
	})
}
