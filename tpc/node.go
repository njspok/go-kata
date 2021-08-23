package tpc

import "fmt"

const (
	CommittedStatus Status = "committed"
)

type Status string

type NodeID int

func NewNode(id NodeID) *Node {
	return &Node{
		id:  id,
		log: []string{},
	}
}

type Node struct {
	id  NodeID
	log []string
}

func (n *Node) ID() NodeID {
	return n.id
}

func (n *Node) Prepare(task *Task) error {
	n.log = append(n.log, fmt.Sprintf("prepare %v", task.ID))
	return nil
}

func (n *Node) Commit(id TaskID) error {
	n.log = append(n.log, fmt.Sprintf("commit %v", id))
	return nil
}

func (n *Node) TaskStatus(id TaskID) Status {
	return CommittedStatus
}

func (n *Node) Log() []string {
	return n.log
}
