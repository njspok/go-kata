package atm

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestATM(t *testing.T) {
	t.Run("create", func(t *testing.T) {
		atm := NewATM()
		require.NotNil(t, atm)
	})
	t.Run("put coin", func(t *testing.T) {
		atm := NewATM()
		require.Zero(t, atm.Total())

		err := atm.Put(1)
		require.NoError(t, err)
		require.Equal(t, Sum(1), atm.Total())

		err = atm.Put(2)
		require.NoError(t, err)
		require.Equal(t, Sum(3), atm.Total())

		err = atm.Put(5)
		require.NoError(t, err)
		require.Equal(t, Sum(8), atm.Total())

		err = atm.Put(10)
		require.NoError(t, err)
		require.Equal(t, Sum(18), atm.Total())

		stat := atm.Cache()
		require.Equal(
			t,
			Cache{
				1:  1,
				2:  1,
				5:  1,
				10: 1,
			},
			stat,
		)
	})
	t.Run("give coins", func(t *testing.T) {
		fixtures := []struct {
			name   string
			load   Cache
			sum    Sum
			result Cache
			err    error
		}{
			{
				name: "give success 1",
				load: Cache{
					1:  3,
					2:  1,
					5:  1,
					10: 2,
				},
				sum: 19,
				result: Cache{
					10: 1,
					5:  1,
					2:  1,
					1:  2,
				},
				err: nil,
			},
			{
				name: "give success 2",
				load: Cache{
					1:  3,
					2:  1,
					5:  1,
					10: 2,
				},
				sum: 20,
				result: Cache{
					10: 2,
				},
				err: nil,
			},
			{
				name: "try give more than in atm",
				load: Cache{
					1:  3,
					2:  1,
					5:  1,
					10: 2,
				},
				sum:    100,
				result: nil,
				err:    ErrNotEnoughCoins,
			},
			{
				name: "no match coins",
				load: Cache{
					2:  1,
					5:  1,
					10: 2,
				},
				sum: 26,
				result: Cache{
					10: 2,
					5:  1,
				},
				err: ErrNotEnoughCoins,
			},
			{
				name: "give all coins",
				load: Cache{
					1:  3,
					2:  5,
					5:  2,
					10: 1,
				},
				sum: 33,
				result: Cache{
					10: 1,
					5:  2,
					2:  5,
					1:  3,
				},
				err: nil,
			},
		}

		for _, f := range fixtures {
			t.Run(f.name, func(t *testing.T) {
				atm := NewATM()
				atm.Load(f.load)
				result, err := atm.Give(f.sum)
				require.ErrorIs(t, err, f.err)
				require.Equal(t, f.result, result)
			})
		}
	})
	t.Run("cache", func(t *testing.T) {
		atm := NewATM()
		atm.Load(Cache{
			1:  9,
			2:  3,
			5:  7,
			10: 11,
		})
		cache, err := atm.Give(28)
		require.NoError(t, err)
		require.Equal(t, Cache{
			10: 2,
			5:  1,
			2:  1,
			1:  1,
		}, cache)
		remains := atm.Cache()
		require.Equal(t, Cache{
			1:  8,
			2:  2,
			5:  6,
			10: 9,
		}, remains)
	})
}
