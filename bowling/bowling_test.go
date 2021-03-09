package bowling

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBowling(t *testing.T) {
	t.Run("create", func(t *testing.T) {
		game := NewBowling()
		require.NotNil(t, game)
	})
	t.Run("rolls", func(t *testing.T) {
		type frames []struct {
			number    uint
			rolls     []uint
			total     uint
			statFrame Frame
		}

		type final struct {
			rolls     []uint
			total     uint
			statFrame FinalFrame
		}

		fixtures := []struct {
			name   string
			frames frames
			final  final
		}{
			{
				name: "all frames opened",
				frames: frames{
					{
						number: 1,
						rolls:  []uint{1, 2},
						total:  3,
						statFrame: Frame{
							Number: 1,
							First:  1,
							Second: 2,
							Score:  3,
							Status: OpenedStatus,
						},
					},
					{
						number: 2,
						rolls:  []uint{3, 2},
						total:  uint(8),
						statFrame: Frame{
							Number: 2,
							First:  3,
							Second: 2,
							Score:  8,
							Status: OpenedStatus,
						},
					},
					{
						number: 3,
						rolls:  []uint{3, 4},
						total:  uint(15),
						statFrame: Frame{
							Number: 3,
							First:  3,
							Second: 4,
							Score:  15,
							Status: OpenedStatus,
						},
					},
					{
						number: 4,
						rolls:  []uint{5, 1},
						total:  uint(21),
						statFrame: Frame{
							Number: 4,
							First:  5,
							Second: 1,
							Score:  21,
							Status: OpenedStatus,
						},
					},
					{
						number: 5,
						rolls:  []uint{3, 4},
						total:  uint(28),
						statFrame: Frame{
							Number: 5,
							First:  3,
							Second: 4,
							Score:  28,
							Status: OpenedStatus,
						},
					},
					{
						number: 6,
						rolls:  []uint{5, 0},
						total:  uint(33),
						statFrame: Frame{
							Number: 6,
							First:  5,
							Second: 0,
							Score:  33,
							Status: OpenedStatus,
						},
					},
					{
						number: 7,
						rolls:  []uint{2, 3},
						total:  uint(38),
						statFrame: Frame{
							Number: 7,
							First:  2,
							Second: 3,
							Score:  38,
							Status: OpenedStatus,
						},
					},
					{
						number: 8,
						rolls:  []uint{4, 3},
						total:  uint(45),
						statFrame: Frame{
							Number: 8,
							First:  4,
							Second: 3,
							Score:  45,
							Status: OpenedStatus,
						},
					},
					{
						number: 9,
						rolls:  []uint{2, 2},
						total:  uint(49),
						statFrame: Frame{
							Number: 9,
							First:  2,
							Second: 2,
							Score:  49,
							Status: OpenedStatus,
						},
					},
				},
				final: final{
					rolls: []uint{7, 2},
					total: 58,
					statFrame: FinalFrame{
						First:  7,
						Second: 2,
						Third:  0,
						Score:  58,
						Status: OpenedStatus,
					},
				},
			},
		}

		for _, fixture := range fixtures {
			t.Run(fixture.name, func(t *testing.T) {
				game := NewBowling()
				stat := NewStat()
				require.Zero(t, game.Total())
				require.Equal(t, stat, game.Stat())

				for _, frame := range fixture.frames {
					t.Run(fmt.Sprintf("frame %v", frame.number), func(t *testing.T) {
						for _, r := range frame.rolls {
							game.Roll(r)
						}
						require.Equal(t, frame.total, game.Total())
						stat.SetFrame(frame.statFrame)
						require.Equal(t, stat, game.Stat())
					})
				}

				// final frame
				for _, r := range fixture.final.rolls {
					game.Roll(r)
				}
				require.Equal(t, fixture.final.total, game.Total())
				require.True(t, game.IsFinished())
				stat.SetFinalFrame(fixture.final.statFrame)
				require.Equal(t, stat, game.Stat())
			})
		}
		//t.Run("all strikes", func(t *testing.T) {
		//	// todo
		//})
		//t.Run("all spares", func(t *testing.T) {
		//	// todo
		//})
		//t.Run("mixed", func(t *testing.T) {
		//	// todo
		//})
	})
}
