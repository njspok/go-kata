package transaction

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransaction(t *testing.T) {
	t.Run("with no operations", func(t *testing.T) {
		tr := NewTransaction()
		ok := tr.Execute()
		require.True(t, ok)
		require.NoError(t, tr.ExecErr())
		require.NoError(t, tr.UndoErr())
	})
	t.Run("success execute one operation", func(t *testing.T) {
		var log []string

		tr := NewTransaction()
		tr.Add(NewOperation(
			func() error {
				log = append(log, "execute one")
				return nil
			}, func() error {
				log = append(log, "undo one")
				return nil
			},
		))

		// check
		ok := tr.Execute()
		require.True(t, ok)
		require.NoError(t, tr.ExecErr())
		require.NoError(t, tr.UndoErr())
		require.Equal(t, []string{
			"execute one",
		}, log)
	})
	t.Run("fail execute one operation", func(t *testing.T) {
		var log []string

		tr := NewTransaction()
		tr.Add(NewOperation(
			func() error {
				log = append(log, "execute one")
				return errors.New("fail one")
			}, func() error {
				log = append(log, "undo one")
				return nil
			},
		))

		// check
		ok := tr.Execute()
		require.False(t, ok)
		require.EqualError(t, tr.ExecErr(), "fail one")
		require.NoError(t, tr.UndoErr())
		require.Equal(t, []string{
			"execute one",
			"undo one",
		}, log)
	})
	t.Run("fail undo one operation", func(t *testing.T) {
		var log []string

		tr := NewTransaction()
		tr.Add(NewOperation(
			func() error {
				log = append(log, "execute one")
				return errors.New("fail execute one")
			}, func() error {
				log = append(log, "undo one")
				return errors.New("fail undo one")
			},
		))

		// check
		ok := tr.Execute()
		require.False(t, ok)
		require.EqualError(t, tr.ExecErr(), "fail execute one")
		require.EqualError(t, tr.UndoErr(), "fail undo one")
		require.Equal(t, []string{
			"execute one",
			"undo one",
		}, log)
	})
	t.Run("success execute", func(t *testing.T) {
		var log []string

		tr := NewTransaction()
		tr.Add(NewOperation(func() error {
			log = append(log, "one")
			return nil
		}, nil))
		tr.Add(NewOperation(func() error {
			log = append(log, "two")
			return nil
		}, nil))

		ok := tr.Execute()
		require.True(t, ok)
		require.NoError(t, tr.UndoErr())
		require.NoError(t, tr.ExecErr())
		require.Equal(t, []string{
			"one",
			"two",
		}, log)
	})
	t.Run("fail execute", func(t *testing.T) {
		var log []string

		tr := NewTransaction()

		// executed and undo
		tr.Add(NewOperation(
			func() error {
				log = append(log, "execute one")
				return nil
			},
			func() error {
				log = append(log, "undo one")
				return nil
			}))

		// executed and undo
		tr.Add(NewOperation(
			func() error {
				log = append(log, "execute two")
				return errors.New("shit happens")
			},
			func() error {
				log = append(log, "undo two")
				return nil
			}))

		// not executed
		tr.Add(NewOperation(
			func() error {
				log = append(log, "execute three")
				return nil
			},
			func() error {
				log = append(log, "undo three")
				return nil
			}))

		// check
		ok := tr.Execute()
		require.False(t, ok)
		require.NoError(t, tr.UndoErr())
		require.EqualError(t, tr.ExecErr(), "shit happens")
		require.Equal(t, []string{
			"execute one",
			"execute two",
			"undo two",
			"undo one",
		}, log)
	})
	t.Run("fail undo", func(t *testing.T) {
		var log []string

		tr := NewTransaction()

		// executed and undo
		tr.Add(NewOperation(
			func() error {
				log = append(log, "execute one")
				return nil
			},
			func() error {
				log = append(log, "undo one")
				return nil
			}))

		// executed and undo
		tr.Add(NewOperation(
			func() error {
				log = append(log, "execute two")
				return nil
			},
			func() error {
				log = append(log, "undo two")
				return errors.New("undo fail")
			}))

		// only executed
		tr.Add(NewOperation(
			func() error {
				log = append(log, "execute three")
				return errors.New("three fail")
			},
			func() error {
				log = append(log, "undo three")
				return nil
			}))

		// check
		ok := tr.Execute()
		require.False(t, ok)
		require.EqualError(t, tr.UndoErr(), "undo fail")
		require.EqualError(t, tr.ExecErr(), "three fail")
		require.Equal(t, []string{
			"execute one",
			"execute two",
			"execute three",
			"undo three",
			"undo two",
		}, log)
	})
}
