package two_phase_commit

import "errors"

var (
	ErrNodesNotExist    = errors.New("nodes not exist")
	ErrNodeAlreadyAdded = errors.New("node already added")
)

type NodeI interface {
	ID() NodeID
	Prepare(TaskI) error
	Commit(TaskID) error
	Abort(TaskID) // todo need return errors?
}

type TaskI interface {
	ID() TaskID
}

func NewTransactionManager() *TransactionManager {
	return &TransactionManager{
		nodes: make(map[NodeID]NodeI),
	}
}

type TransactionManager struct {
	nodes map[NodeID]NodeI
}

func (m *TransactionManager) Run(task TaskI) error {
	if m.withoutNodes() {
		return ErrNodesNotExist
	}

	// todo many errors
	var lastErr error
	var prepared []NodeI
	for _, node := range m.nodes {
		err := node.Prepare(task)
		if err != nil {
			lastErr = err
			continue
		}
		prepared = append(prepared, node)
	}

	if lastErr != nil {
		for _, node := range prepared {
			node.Abort(task.ID())
		}
		return lastErr
	}

	for _, node := range m.nodes {
		err := node.Commit(task.ID())
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *TransactionManager) Add(node NodeI) error {
	if _, exist := m.nodes[node.ID()]; exist {
		return ErrNodeAlreadyAdded
	}
	m.nodes[node.ID()] = node
	return nil
}

func (m *TransactionManager) withoutNodes() bool {
	return len(m.nodes) == 0
}
