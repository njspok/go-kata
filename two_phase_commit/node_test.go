package two_phase_commit

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNode_Commit(t *testing.T) {
	t.Run("task not found", func(t *testing.T) {
		node := NewNode(111)
		require.ErrorIs(t, node.Commit(2222), ErrTaskNotFound)
	})
	t.Run("task not prepare success", func(t *testing.T) {
		node := NewNode(111)
		node.SetPrepareErr(errors.New("shit happens"))

		err := node.Prepare(NewTask(999))
		require.EqualError(t, err, "shit happens")

		err = node.Commit(999)
		require.ErrorIs(t, err, ErrTaskMustPrepareSuccessStatus)
	})
}

func TestNode_TaskStatus(t *testing.T) {
	t.Run("task not found", func(t *testing.T) {
		node := NewNode(111)
		status, err := node.TaskStatus(2222)
		require.ErrorIs(t, err, ErrTaskNotFound)
		require.Equal(t, NoneStatus, status)
	})
}
