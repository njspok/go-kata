package two_phase_commit

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNode_Commit(t *testing.T) {
	t.Run("task not found", func(t *testing.T) {
		node := NewNode(111)
		require.ErrorIs(t, node.Commit(2222), ErrTaskNotFound)
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
