package two_phase_commit

import "fmt"

const (
	CommittedStatus     Status = "committed"
	PrepareFailedStatus Status = "prepare-failed"
)

type Status string

type NodeID int

func NewNode(id NodeID) *Node {
	return &Node{
		id:         id,
		log:        []string{},
		prepareErr: nil,
	}
}

type Node struct {
	id         NodeID
	log        []string
	prepareErr error
}

func (n *Node) ID() NodeID {
	return n.id
}

func (n *Node) Prepare(task TaskI) error {
	if n.prepareErr != nil {
		n.addToLog("prepare %v failed", task.ID())
		return n.prepareErr
	}

	n.addToLog("prepare %v", task.ID())
	return nil
}

func (n *Node) Commit(id TaskID) error {
	n.addToLog("commit %v", id)
	return nil
}

func (n *Node) TaskStatus(id TaskID) Status {
	return CommittedStatus
}

func (n *Node) Log() []string {
	return n.log
}

func (n *Node) SetPrepareErr(err error) {
	n.prepareErr = err
}

func (n *Node) addToLog(s string, a ...interface{}) {
	n.log = append(n.log, fmt.Sprintf(s, a...))
}
