package two_phase_commit

import (
	"errors"
	"fmt"
)

const (
	NoneStatus             Status = "none"
	CommittedSuccessStatus Status = "committed-success"
	CommitFailedStatus     Status = "commit-failed"
	PrepareSuccessStatus   Status = "prepare-success"
	PrepareFailedStatus    Status = "prepare-failed"
	AbortSuccessStatus     Status = "abort-success"
)

var (
	ErrTaskAlreadyExist             = errors.New("task already exist")
	ErrTaskNotFound                 = errors.New("task not found")
	ErrTaskMustPrepareSuccessStatus = errors.New("task must prepare success status")
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

func (n *Node) Abort(id TaskID) {
	// todo return error not found

	status := n.task[id]
	if status != PrepareSuccessStatus {
		// todo return error
		return
	}

	n.setTaskStatus(id, AbortSuccessStatus)
}

func (n *Node) Prepare(task TaskI) error {
	if _, exist := n.task[task.ID()]; exist {
		return ErrTaskAlreadyExist
	}

	if n.prepareErr != nil {
		n.setTaskStatus(task.ID(), PrepareFailedStatus)
		return n.prepareErr
	}

	n.setTaskStatus(task.ID(), PrepareSuccessStatus)
	return nil
}

func (n *Node) Commit(id TaskID) error {
	status, exist := n.task[id]
	if !exist {
		return ErrTaskNotFound
	}

	if status != PrepareSuccessStatus {
		return ErrTaskMustPrepareSuccessStatus
	}

	if n.commitErr != nil {
		n.setTaskStatus(id, CommitFailedStatus)
		return n.commitErr
	}

	n.setTaskStatus(id, CommittedSuccessStatus)
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

func (n *Node) SetCommitErr(err error) {
	n.commitErr = err
}

func (n *Node) addToLog(s string, a ...interface{}) {
	n.log = append(n.log, fmt.Sprintf(s, a...))
}

func (n *Node) setTaskStatus(id TaskID, status Status) {
	// todo make map
	switch status {
	case PrepareSuccessStatus:
		n.addToLog("prepare %v success", id)
	case PrepareFailedStatus:
		n.addToLog("prepare %v failed", id)
	case CommitFailedStatus:
		n.addToLog("commit %v failed", id)
	case CommittedSuccessStatus:
		n.addToLog("commit %v success", id)
	case AbortSuccessStatus:
		n.addToLog("abort %v success", id)
	default:
		panic("unknown status")
	}

	n.task[id] = status
}
