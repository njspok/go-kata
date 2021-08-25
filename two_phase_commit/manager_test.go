package two_phase_commit

import (
	"errors"
	"fmt"
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

			manager, node := makeManagerAndNode(t)

			// run task
			task := NewTask(1)
			err = manager.Run(task)
			require.NoError(t, err)

			// check
			status, err := node.TaskStatus(task.ID())
			require.NoError(t, err)
			require.Equal(t, CommittedSuccessStatus, status)
			require.Equal(t, []string{
				"prepare 1 success",
				"commit 1 success",
			}, node.Log())
		})
		t.Run("run twice", func(t *testing.T) {
			manager, _ := makeManagerAndNode(t)

			// create task
			task := NewTask(1)

			// run
			require.NoError(t, manager.Run(task))
			require.ErrorIs(t, manager.Run(task), ErrTaskAlreadyExist)
		})
		t.Run("prepare failed", func(t *testing.T) {
			var err error
			ErrSomeNodePrepareError := errors.New("some node prepare error")

			manager, node := makeManagerAndNode(t)

			// broken node
			node.SetPrepareErr(ErrSomeNodePrepareError)

			// run task
			task := NewTask(1)
			err = manager.Run(task)
			require.ErrorIs(t, err, ErrSomeNodePrepareError)

			// check
			status, err := node.TaskStatus(task.ID())
			require.NoError(t, err)
			require.Equal(t, PrepareFailedStatus, status)
			require.Equal(t, []string{
				"prepare 1 failed",
			}, node.Log())
		})
		t.Run("commit failed", func(t *testing.T) {
			var err error
			ErrSomeNodeCommitError := errors.New("some node commit error")

			manager, node := makeManagerAndNode(t)

			// broken commit node
			node.SetCommitErr(ErrSomeNodeCommitError)

			// run task
			task := NewTask(1)
			err = manager.Run(task)
			require.ErrorIs(t, err, ErrSomeNodeCommitError)

			// check
			status, err := node.TaskStatus(task.ID())
			require.NoError(t, err)
			require.Equal(t, CommitFailedStatus, status)
			require.Equal(t, []string{
				"prepare 1 success",
				"commit 1 failed",
			}, node.Log())
		})
	})
	t.Run("multiple nodes", func(t *testing.T) {
		t.Run("success", func(t *testing.T) {
			manager, nodes := makeManagerAndNodes(t)

			// run task
			task := NewTask(1)
			err := manager.Run(task)
			require.NoError(t, err)

			// check
			for i, node := range nodes {
				t.Run(fmt.Sprintf("check node %v", i), func(t *testing.T) {
					status, err := node.TaskStatus(task.ID())
					require.NoError(t, err)
					require.Equal(t, CommittedSuccessStatus, status)
					require.Equal(t, []string{
						"prepare 1 success",
						"commit 1 success",
					}, node.Log())
				})
			}
		})
		t.Run("prepare failed", func(t *testing.T) {
			t.Run("failed 1 node", func(t *testing.T) {
				manager, nodes := makeManagerAndNodes(t)

				// broken node
				nodes[300].SetPrepareErr(errors.New("shit happens"))

				// run task
				task := NewTask(1)
				err := manager.Run(task)
				require.EqualError(t, err, "shit happens")

				// check

				checkNodes(t, func(t *testing.T, node *Node) {
					status, err := node.TaskStatus(task.ID())
					require.NoError(t, err)
					require.Equal(t, AbortSuccessStatus, status)
					require.Equal(t, []string{
						"prepare 1 success",
						"abort 1 success",
					}, node.Log())
				}, nodes, 100, 200)

				checkNodes(t, func(t *testing.T, node *Node) {
					status, err := node.TaskStatus(task.ID())
					require.NoError(t, err)
					require.Equal(t, PrepareFailedStatus, status)
					require.Equal(t, []string{
						"prepare 1 failed",
					}, node.Log())
				}, nodes, 300)
			})
			t.Run("failed all nodes", func(t *testing.T) {
				manager, nodes := makeManagerAndNodes(t)

				// broken all nodes
				for _, node := range nodes {
					node.SetPrepareErr(errors.New("shit happens"))
				}

				// run task
				task := NewTask(1)
				err := manager.Run(task)
				require.EqualError(t, err, "shit happens")

				// check

				checkNodes(t, func(t *testing.T, node *Node) {
					status, err := node.TaskStatus(task.ID())
					require.NoError(t, err)
					require.Equal(t, PrepareFailedStatus, status)
					require.Equal(t, []string{
						"prepare 1 failed",
					}, node.Log())
				}, nodes, 100, 200, 300)
			})
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

func checkNodes(
	t *testing.T,
	f func(t *testing.T, node *Node),
	nodes map[NodeID]*Node,
	ids ...NodeID,
) {
	for _, id := range ids {
		node := nodes[id]
		t.Run(fmt.Sprintf("check node %v", id), func(t *testing.T) {
			f(t, node)
		})
	}
}

func makeManagerAndNode(t *testing.T) (*TransactionManager, *Node) {
	manager := NewTransactionManager()
	require.NotNil(t, manager)

	node := NewNode(100)
	require.NotNil(t, node)

	require.NoError(t, manager.Add(node))

	return manager, node
}

func makeManagerAndNodes(t *testing.T) (*TransactionManager, map[NodeID]*Node) {
	manager := NewTransactionManager()
	require.NotNil(t, manager)

	// add node
	node100 := NewNode(100)
	require.NoError(t, manager.Add(node100))

	// add node
	node200 := NewNode(200)
	require.NoError(t, manager.Add(node200))

	// add node
	node300 := NewNode(300)
	require.NoError(t, manager.Add(node300))

	return manager, map[NodeID]*Node{
		100: node100,
		200: node200,
		300: node300,
	}
}
