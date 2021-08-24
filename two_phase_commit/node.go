package two_phase_commit

import (
	"errors"
	"fmt"
)

const (
	CommittedSuccessStatus Status = "committed-success"
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
		task:       make(map[TaskID]Status),
	}
}

type Node struct {
	id         NodeID
	log        []string
	prepareErr error
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
	n.addToLog("prepare %v", task.ID())
	return nil
}

func (n *Node) Commit(id TaskID) error {
	if _, exist := n.task[id]; exist {
		n.task[id] = CommittedSuccessStatus
		n.addToLog("commit %v", id)
		return nil
	}

	return ErrTaskNotFound
}

func (n *Node) TaskStatus(id TaskID) Status {
	status, exist := n.task[id]
	if !exist {
		return "" // todo repalce interface
	}

	return status
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
