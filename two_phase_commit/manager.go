package two_phase_commit

import "errors"

var (
	ErrNodePrepareFailed = errors.New("node prepare failed")
	ErrNodesNotExist     = errors.New("nodes not exist")
	ErrNodeAlreadyAdded  = errors.New("node already added")
)

type NodeI interface {
	ID() NodeID
	Prepare(TaskI) error
	Commit(TaskID) error
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

	for _, node := range m.nodes {
		err := node.Prepare(task)
		if err != nil {
			return err
		}
	}

	for _, node := range m.nodes {
		// todo process errors
		_ = node.Commit(task.ID())
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
