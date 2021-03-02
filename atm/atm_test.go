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
		require.Equal(t, uint(1), atm.Total())

		err = atm.Put(2)
		require.NoError(t, err)
		require.Equal(t, uint(3), atm.Total())

		err = atm.Put(5)
		require.NoError(t, err)
		require.Equal(t, uint(8), atm.Total())

		err = atm.Put(10)
		require.NoError(t, err)
		require.Equal(t, uint(18), atm.Total())

		stat := atm.Stat()
		require.Equal(
			t,
			Stat{
				1:  1,
				2:  1,
				5:  1,
				10: 1,
			},
			stat,
		)
	})
	t.Run("give coins", func(t *testing.T) {
		// todo
	})
	t.Run("cant give coins", func(t *testing.T) {
		// todo
	})
}
