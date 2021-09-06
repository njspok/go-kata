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
	Abort(TaskID) error // todo need return errors?
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

	if err := m.prepare(task); err != nil {
		return err
	}

	if err := m.commit(task); err != nil {
		return err
	}

	return nil
}

func (m *TransactionManager) Add(node NodeI) error {
	if m.nodeExist(node.ID()) {
		return ErrNodeAlreadyAdded
	}
	m.setNode(node)
	return nil
}

func (m *TransactionManager) nodeExist(id NodeID) bool {
	_, exist := m.nodes[id]
	return exist
}

func (m *TransactionManager) setNode(node NodeI) {
	m.nodes[node.ID()] = node
}

func (m *TransactionManager) withoutNodes() bool {
	return len(m.nodes) == 0
}

func (m *TransactionManager) prepare(task TaskI) error {
	var errs ErrorsList
	var prepared []NodeI

	for _, node := range m.nodes {
		err := node.Prepare(task)
		if err != nil {
			errs.Add(err)
			continue
		}
		prepared = append(prepared, node)
	}

	if !errs.Empty() {
		err := m.abort(task, prepared)
		if err != nil {
			errs.Add(err)
		}
		return errs
	}

	return nil
}

func (m *TransactionManager) commit(task TaskI) error {
	for _, node := range m.nodes {
		err := node.Commit(task.ID())
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *TransactionManager) abort(task TaskI, nodes []NodeI) error {
	for _, node := range nodes {
		if err := node.Abort(task.ID()); err != nil {
			return err
		}
	}
	return nil
}
