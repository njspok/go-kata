package tpc

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransactionManager_Run(t *testing.T) {
	t.Run("without nodes", func(t *testing.T) {
		manager := NewTransactionManager()
		err := manager.Run(NewTask(1))
		require.ErrorIs(t, err, ErrNodesNotExist)
	})
	t.Run("one node", func(t *testing.T) {
		t.Run("success", func(t *testing.T) {
			var err error

			manager := NewTransactionManager()

			// add node
			node := NewNode(100)
			err = manager.Add(node)
			require.NoError(t, err)

			// run task
			task := NewTask(1)
			err = manager.Run(task)
			require.NoError(t, err)

			// check

			status := node.TaskStatus(task.ID())
			require.Equal(t, CommittedStatus, status)

			log := node.Log()
			require.Equal(t, []string{
				"prepare 1",
				"commit 1",
			}, log)
		})
		t.Run("prepare failed", func(t *testing.T) {
			// todo
		})
		t.Run("commit failed", func(t *testing.T) {
			// todo
		})
	})
	t.Run("multiple nodes", func(t *testing.T) {
		t.Run("success", func(t *testing.T) {
			var err error

			manager := NewTransactionManager()

			// add node
			node100 := NewNode(100)
			err = manager.Add(node100)
			require.NoError(t, err)

			// add node
			node200 := NewNode(200)
			err = manager.Add(node200)
			require.NoError(t, err)

			// add node
			node300 := NewNode(300)
			err = manager.Add(node300)
			require.NoError(t, err)

			// run task
			task := NewTask(1)
			err = manager.Run(task)
			require.NoError(t, err)

			// check

			// check node
			require.Equal(t, CommittedStatus, node100.TaskStatus(task.ID()))
			require.Equal(t, []string{
				"prepare 1",
				"commit 1",
			}, node100.Log())

			// check node
			require.Equal(t, CommittedStatus, node200.TaskStatus(task.ID()))
			require.Equal(t, []string{
				"prepare 1",
				"commit 1",
			}, node200.Log())

			// check node 1
			require.Equal(t, CommittedStatus, node300.TaskStatus(task.ID()))
			require.Equal(t, []string{
				"prepare 1",
				"commit 1",
			}, node300.Log())
		})
		t.Run("prepare failed", func(t *testing.T) {
			// todo
		})
		t.Run("commit failed", func(t *testing.T) {
			// todo
		})
	})
}

func TestTransactionManager_Add(t *testing.T) {
	t.Run("try add same node again", func(t *testing.T) {
		manager := NewTransactionManager()

		node := NewNode(100)

		require.NoError(t, manager.Add(node))
		require.ErrorIs(t, manager.Add(node), ErrNodeAlreadyAdded)
	})
	t.Run("add different nodes", func(t *testing.T) {
		manager := NewTransactionManager()

		require.NoError(t, manager.Add(NewNode(100)))
		require.NoError(t, manager.Add(NewNode(200)))
		require.NoError(t, manager.Add(NewNode(300)))
	})
}
