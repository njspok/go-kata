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
			stat := NewStat()
			require.Zero(t, game.Total())
			require.Equal(t, stat, game.Stat())

			// 1
			game.Roll(1)
			game.Roll(2)
			require.Equal(t, uint(3), game.Total())
			stat.SetFrame(Frame{
				Number: 1,
				First:  1,
				Second: 2,
				Score:  3,
				Status: OpenedStatus,
			})
			require.Equal(t, stat, game.Stat())

			// 2
			game.Roll(3)
			game.Roll(2)
			require.Equal(t, uint(8), game.Total())
			stat.SetFrame(Frame{
				Number: 2,
				First:  3,
				Second: 2,
				Score:  8,
				Status: OpenedStatus,
			})
			require.Equal(t, stat, game.Stat())

			// 3
			game.Roll(3)
			game.Roll(4)
			require.Equal(t, uint(15), game.Total())
			stat.SetFrame(Frame{
				Number: 3,
				First:  3,
				Second: 4,
				Score:  15,
				Status: OpenedStatus,
			})
			require.Equal(t, stat, game.Stat())

			// 4
			game.Roll(5)
			game.Roll(1)
			require.Equal(t, uint(21), game.Total())
			stat.SetFrame(Frame{
				Number: 4,
				First:  5,
				Second: 1,
				Score:  21,
				Status: OpenedStatus,
			})
			require.Equal(t, stat, game.Stat())

			// 5
			game.Roll(3)
			game.Roll(4)
			require.Equal(t, uint(28), game.Total())
			stat.SetFrame(Frame{
				Number: 5,
				First:  3,
				Second: 4,
				Score:  28,
				Status: OpenedStatus,
			})
			require.Equal(t, stat, game.Stat())

			// 6
			game.Roll(5)
			game.Roll(0)
			require.Equal(t, uint(33), game.Total())
			stat.SetFrame(Frame{
				Number: 6,
				First:  5,
				Second: 0,
				Score:  33,
				Status: OpenedStatus,
			})
			require.Equal(t, stat, game.Stat())

			// 7
			game.Roll(2)
			game.Roll(3)
			require.Equal(t, uint(38), game.Total())
			stat.SetFrame(Frame{
				Number: 7,
				First:  2,
				Second: 3,
				Score:  38,
				Status: OpenedStatus,
			})
			require.Equal(t, stat, game.Stat())

			// 8
			game.Roll(4)
			game.Roll(3)
			require.Equal(t, uint(45), game.Total())
			stat.SetFrame(Frame{
				Number: 8,
				First:  4,
				Second: 3,
				Score:  45,
				Status: OpenedStatus,
			})
			require.Equal(t, stat, game.Stat())

			// 9
			game.Roll(2)
			game.Roll(2)
			require.Equal(t, uint(49), game.Total())
			stat.SetFrame(Frame{
				Number: 9,
				First:  2,
				Second: 2,
				Score:  49,
				Status: OpenedStatus,
			})
			require.Equal(t, stat, game.Stat())

			// 10
			game.Roll(7)
			game.Roll(2)
			err := game.Roll(5)
			require.ErrorIs(t, err, ErrGameFinished)
			require.Equal(t, uint(58), game.Total())
			require.True(t, game.IsFinished())
			stat.SetFinalFrame(FinalFrame{
				First:  7,
				Second: 2,
				Third:  0,
				Score:  58,
				Status: OpenedStatus,
			})
			require.Equal(t, stat, game.Stat())
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
