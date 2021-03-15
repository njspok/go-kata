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
		type frame struct {
			number  uint
			rolls   []uint
			first   uint
			second  uint
			score   uint
			status  Status
			bonuses [9]uint
		}

		type frames []frame

		type final struct {
			number  uint
			rolls   []uint
			first   uint
			second  uint
			third   uint
			score   uint
			status  Status
			bonuses [10]uint
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
						number:  1,
						rolls:   []uint{1, 2},
						first:   1,
						second:  2,
						score:   3,
						status:  OpenedStatus,
						bonuses: [9]uint{0, 0, 0, 0, 0, 0, 0, 0, 0},
					},
					{
						number:  2,
						rolls:   []uint{3, 2},
						first:   3,
						second:  2,
						score:   8,
						status:  OpenedStatus,
						bonuses: [9]uint{0, 0, 0, 0, 0, 0, 0, 0, 0},
					},
					{
						number:  3,
						rolls:   []uint{3, 4},
						first:   3,
						second:  4,
						score:   15,
						status:  OpenedStatus,
						bonuses: [9]uint{0, 0, 0, 0, 0, 0, 0, 0, 0},
					},
					{
						number:  4,
						rolls:   []uint{5, 1},
						first:   5,
						second:  1,
						score:   21,
						status:  OpenedStatus,
						bonuses: [9]uint{0, 0, 0, 0, 0, 0, 0, 0, 0},
					},
					{
						number:  5,
						rolls:   []uint{3, 4},
						first:   3,
						second:  4,
						score:   28,
						status:  OpenedStatus,
						bonuses: [9]uint{0, 0, 0, 0, 0, 0, 0, 0, 0},
					},
					{
						number:  6,
						rolls:   []uint{5, 0},
						first:   5,
						second:  0,
						score:   33,
						status:  OpenedStatus,
						bonuses: [9]uint{0, 0, 0, 0, 0, 0, 0, 0, 0},
					},
					{
						number:  7,
						rolls:   []uint{2, 3},
						first:   2,
						second:  3,
						score:   38,
						status:  OpenedStatus,
						bonuses: [9]uint{0, 0, 0, 0, 0, 0, 0, 0, 0},
					},
					{
						number:  8,
						rolls:   []uint{4, 3},
						first:   4,
						second:  3,
						score:   45,
						status:  OpenedStatus,
						bonuses: [9]uint{0, 0, 0, 0, 0, 0, 0, 0, 0},
					},
					{
						number:  9,
						rolls:   []uint{2, 2},
						first:   2,
						second:  2,
						score:   49,
						status:  OpenedStatus,
						bonuses: [9]uint{0, 0, 0, 0, 0, 0, 0, 0, 0},
					},
				},
				final: final{
					number:  10,
					rolls:   []uint{7, 2},
					first:   7,
					second:  2,
					third:   0,
					score:   58,
					status:  FinalStatus,
					bonuses: [10]uint{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				},
			},
			{
				name: "mixed frame status",
				frames: frames{
					{
						number:  1,
						rolls:   []uint{10},
						first:   10,
						second:  0,
						score:   10,
						status:  StrikeStatus,
						bonuses: [9]uint{0, 0, 0, 0, 0, 0, 0, 0, 0},
					},
					{
						number:  2,
						rolls:   []uint{10},
						first:   10,
						second:  0,
						score:   20,
						status:  StrikeStatus,
						bonuses: [9]uint{0, 0, 0, 0, 0, 0, 0, 0, 0},
					},
					{
						number:  3,
						rolls:   []uint{6, 4},
						first:   6,
						second:  4,
						score:   30,
						status:  SpareStatus,
						bonuses: [9]uint{0, 0, 0, 0, 0, 0, 0, 0, 0},
					},
					{
						number:  4,
						rolls:   []uint{0, 0},
						first:   0,
						second:  0,
						score:   30,
						status:  OpenedStatus,
						bonuses: [9]uint{0, 0, 0, 0, 0, 0, 0, 0, 0},
					},
					{
						number:  5,
						rolls:   []uint{5, 4},
						first:   5,
						second:  4,
						score:   39,
						status:  OpenedStatus,
						bonuses: [9]uint{0, 0, 0, 0, 0, 0, 0, 0, 0},
					},
					{
						number:  6,
						rolls:   []uint{2, 1},
						first:   2,
						second:  1,
						score:   42,
						status:  OpenedStatus,
						bonuses: [9]uint{0, 0, 0, 0, 0, 0, 0, 0, 0},
					},
					{
						number:  7,
						rolls:   []uint{9, 1},
						first:   9,
						second:  1,
						score:   52,
						status:  SpareStatus,
						bonuses: [9]uint{0, 0, 0, 0, 0, 0, 0, 0, 0},
					},
					{
						number:  8,
						rolls:   []uint{10},
						first:   10,
						second:  0,
						score:   62,
						status:  StrikeStatus,
						bonuses: [9]uint{0, 0, 0, 0, 0, 0, 0, 0, 0},
					},
					{
						number:  9,
						rolls:   []uint{3, 5},
						first:   3,
						second:  5,
						score:   70,
						status:  OpenedStatus,
						bonuses: [9]uint{0, 0, 0, 0, 0, 0, 0, 0, 0},
					},
				},
				final: final{
					number:  10,
					rolls:   []uint{9, 1, 10},
					first:   9,
					second:  1,
					third:   10,
					score:   90,
					status:  FinalStatus,
					bonuses: [10]uint{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				},
			},
			{
				name: "all strikes",
				frames: frames{
					{
						number:  1,
						rolls:   []uint{10},
						first:   10,
						second:  0,
						score:   10,
						status:  StrikeStatus,
						bonuses: [9]uint{0, 0, 0, 0, 0, 0, 0, 0, 0},
					},
					{
						number:  2,
						rolls:   []uint{10},
						first:   10,
						second:  0,
						score:   20,
						status:  StrikeStatus,
						bonuses: [9]uint{0, 0, 0, 0, 0, 0, 0, 0, 0},
					},
					{
						number:  3,
						rolls:   []uint{10},
						first:   10,
						second:  0,
						score:   30,
						status:  StrikeStatus,
						bonuses: [9]uint{0, 0, 0, 0, 0, 0, 0, 0, 0},
					},
					{
						number:  4,
						rolls:   []uint{10},
						first:   10,
						second:  0,
						score:   40,
						status:  StrikeStatus,
						bonuses: [9]uint{0, 0, 0, 0, 0, 0, 0, 0, 0},
					},
					{
						number:  5,
						rolls:   []uint{10},
						first:   10,
						second:  0,
						score:   50,
						status:  StrikeStatus,
						bonuses: [9]uint{0, 0, 0, 0, 0, 0, 0, 0, 0},
					},
					{
						number:  6,
						rolls:   []uint{10},
						first:   10,
						second:  0,
						score:   60,
						status:  StrikeStatus,
						bonuses: [9]uint{0, 0, 0, 0, 0, 0, 0, 0, 0},
					},
					{
						number:  7,
						rolls:   []uint{10},
						first:   10,
						second:  0,
						score:   70,
						status:  StrikeStatus,
						bonuses: [9]uint{0, 0, 0, 0, 0, 0, 0, 0, 0},
					},
					{
						number:  8,
						rolls:   []uint{10},
						first:   10,
						second:  0,
						score:   80,
						status:  StrikeStatus,
						bonuses: [9]uint{0, 0, 0, 0, 0, 0, 0, 0, 0},
					},
					{
						number:  9,
						rolls:   []uint{10},
						first:   10,
						second:  0,
						score:   90,
						status:  StrikeStatus,
						bonuses: [9]uint{0, 0, 0, 0, 0, 0, 0, 0, 0},
					},
				},
				final: final{
					number:  10,
					rolls:   []uint{10, 10, 10},
					first:   10,
					second:  10,
					third:   10,
					score:   120,
					status:  FinalStatus,
					bonuses: [10]uint{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				},
			},
			{
				name: "zero",
				frames: frames{
					{
						number:  1,
						rolls:   []uint{0, 0},
						first:   0,
						second:  0,
						score:   0,
						status:  OpenedStatus,
						bonuses: [9]uint{0, 0, 0, 0, 0, 0, 0, 0, 0},
					},
					{
						number:  2,
						rolls:   []uint{0, 0},
						first:   0,
						second:  0,
						score:   0,
						status:  OpenedStatus,
						bonuses: [9]uint{0, 0, 0, 0, 0, 0, 0, 0, 0},
					},
					{
						number:  3,
						rolls:   []uint{0, 0},
						first:   0,
						second:  0,
						score:   0,
						status:  OpenedStatus,
						bonuses: [9]uint{0, 0, 0, 0, 0, 0, 0, 0, 0},
					},
					{
						number:  4,
						rolls:   []uint{0, 0},
						first:   0,
						second:  0,
						score:   0,
						status:  OpenedStatus,
						bonuses: [9]uint{0, 0, 0, 0, 0, 0, 0, 0, 0},
					},
					{
						number:  5,
						rolls:   []uint{0, 0},
						first:   0,
						second:  0,
						score:   0,
						status:  OpenedStatus,
						bonuses: [9]uint{0, 0, 0, 0, 0, 0, 0, 0, 0},
					},
					{
						number:  6,
						rolls:   []uint{0, 0},
						first:   0,
						second:  0,
						score:   0,
						status:  OpenedStatus,
						bonuses: [9]uint{0, 0, 0, 0, 0, 0, 0, 0, 0},
					},
					{
						number:  7,
						rolls:   []uint{0, 0},
						first:   0,
						second:  0,
						score:   0,
						status:  OpenedStatus,
						bonuses: [9]uint{0, 0, 0, 0, 0, 0, 0, 0, 0},
					},
					{
						number:  8,
						rolls:   []uint{0, 0},
						first:   0,
						second:  0,
						score:   0,
						status:  OpenedStatus,
						bonuses: [9]uint{0, 0, 0, 0, 0, 0, 0, 0, 0},
					},
					{
						number:  9,
						rolls:   []uint{0, 0},
						first:   0,
						second:  0,
						score:   0,
						status:  OpenedStatus,
						bonuses: [9]uint{0, 0, 0, 0, 0, 0, 0, 0, 0},
					},
				},
				final: final{
					number:  10,
					rolls:   []uint{0, 0},
					first:   0,
					second:  0,
					third:   0,
					score:   0,
					status:  FinalStatus,
					bonuses: [10]uint{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				},
			},
		}

		for _, fixture := range fixtures {
			t.Run(fixture.name, func(t *testing.T) {
				game := NewBowling()
				require.Zero(t, game.Total())

				for _, ff := range fixture.frames {
					t.Run(fmt.Sprintf("frame %v", ff.number), func(t *testing.T) {
						for _, r := range ff.rolls {
							game.Roll(r)
						}
						require.Equal(t, ff.score, game.Total())

						frame := game.Frames().Frame(ff.number).(*Frame)
						require.Equal(t, ff.number, frame.Number())
						require.Equal(t, ff.first, frame.First())
						require.Equal(t, ff.second, frame.Second())
						require.Equal(t, ff.score, frame.Score())
						require.Equal(t, ff.status, frame.Status())

						// check bonuses
						for i := 1; i <= len(ff.bonuses); i++ {
							frame := game.Frames().Frame(uint(i)).(*Frame)
							require.Equal(t, ff.bonuses[i-1], frame.Bonuses())
						}
					})
				}

				// final frame
				for _, r := range fixture.final.rolls {
					game.Roll(r)
				}
				require.Equal(t, fixture.final.score, game.Total())
				require.True(t, game.IsFinished())

				frame := game.Frames().Frame(10).(*FinalFrame)
				require.Equal(t, fixture.final.number, frame.Number())
				require.Equal(t, fixture.final.first, frame.First())
				require.Equal(t, fixture.final.second, frame.Second())
				require.Equal(t, fixture.final.third, frame.Third())
				require.Equal(t, fixture.final.score, frame.Score())
				require.Equal(t, fixture.final.status, frame.Status())

				// check bonuses
				for i := 1; i <= len(fixture.final.bonuses); i++ {
					frame := game.Frames().Frame(uint(i)).(interface{ Bonuses() uint })
					require.Equal(t, fixture.final.bonuses[i-1], frame.Bonuses())
				}
			})
		}
	})
}
