package tpc

import "errors"

var (
	ErrNodesNotExist    = errors.New("nodes not exist")
	ErrNodeAlreadyAdded = errors.New("node already added")
)

type NodeI interface {
	ID() NodeID
	Prepare(*Task) error
	Commit(TaskID) error
}

func NewTransactionManager() *TransactionManager {
	return &TransactionManager{
		nodes: make(map[NodeID]NodeI),
	}
}

type TransactionManager struct {
	nodes map[NodeID]NodeI
}

func (m *TransactionManager) Run(task *Task) error {
	if m.withoutNodes() {
		return ErrNodesNotExist
	}

	for _, node := range m.nodes {
		// todo process errors
		_ = node.Prepare(task)

		// todo process errors
		_ = node.Commit(task.ID)
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
