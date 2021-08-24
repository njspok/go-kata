package two_phase_commit

import (
	"errors"
	"fmt"
)

const (
	NoneStatus             Status = "none"
	CommittedSuccessStatus Status = "committed-success"
	CommitFailedStatus     Status = "commit-failed"
	PrepareFailedStatus    Status = "prepare-failed"
	PrepareSuccessStatus   Status = "prepare-success"
)

var (
	ErrTaskAlreadyExist = errors.New("task already exist")
	ErrTaskNotFound     = errors.New("task not found")
)

type Status string

type NodeID int

func NewNode(id NodeID) *Node {
	return &Node{
		id:         id,
		log:        []string{},
		prepareErr: nil,
		commitErr:  nil,
		task:       make(map[TaskID]Status),
	}
}

type Node struct {
	id         NodeID
	log        []string
	prepareErr error
	commitErr  error
	task       map[TaskID]Status
}

func (n *Node) ID() NodeID {
	return n.id
}

func (n *Node) Prepare(task TaskI) error {
	if _, exist := n.task[task.ID()]; exist {
		return ErrTaskAlreadyExist
	}

	if n.prepareErr != nil {
		n.task[task.ID()] = PrepareFailedStatus
		n.addToLog("prepare %v failed", task.ID())
		return n.prepareErr
	}

	n.task[task.ID()] = PrepareSuccessStatus
	n.addToLog("prepare %v success", task.ID())
	return nil
}

func (n *Node) Commit(id TaskID) error {
	if _, exist := n.task[id]; !exist {
		return ErrTaskNotFound
	}

	if n.commitErr != nil {
		n.task[id] = CommitFailedStatus
		n.addToLog("commit %v failed", id)
		return n.commitErr
	}

	n.task[id] = CommittedSuccessStatus
	n.addToLog("commit %v success", id)
	return nil
}

func (n *Node) TaskStatus(id TaskID) (Status, error) {
	status, exist := n.task[id]
	if !exist {
		return NoneStatus, ErrTaskNotFound
	}

	return status, nil
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

func (n *Node) SetCommitErr(err error) {
	n.commitErr = err
}
