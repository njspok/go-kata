package two_phase_commit

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestErrorsList_Error(t *testing.T) {
	t.Run("empty errors", func(t *testing.T) {
		errs := ErrorsList{}
		require.EqualError(t, errs, "list without errors")
	})
	t.Run("multiple errors", func(t *testing.T) {
		errs := ErrorsList{}
		errs.Add(errors.New("some one error"))
		errs.Add(errors.New("some two error"))
		errs.Add(errors.New("some three error"))
		require.EqualError(t, errs, "list of 3 errors")
	})
}

func TestErrorsList_Add(t *testing.T) {
	var list ErrorsList

	list.Add(errors.New("one"))
	list.Add(errors.New("two"))
	list.Add(errors.New("three"))

	require.Len(t, list, 3)
	require.Equal(t, 3, list.Count())

	require.EqualError(t, list[0], "one")
	require.EqualError(t, list[1], "two")
	require.EqualError(t, list[2], "three")
}
