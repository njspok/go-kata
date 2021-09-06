package two_phase_commit

import (
	"errors"
	"fmt"
)

const (
	NoneStatus           Status = "none"
	PrepareSuccessStatus Status = "prepare-success"
	PrepareFailedStatus  Status = "prepare-failed"
	CommittedStatus      Status = "committed-success"
	AbortStatus          Status = "abort-success"
)

var (
	ErrTaskAlreadyExist = errors.New("task already exist")
	ErrTaskNotFound     = errors.New("task not found")
	ErrTaskFinished     = errors.New("task finished")
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

func (n *Node) Abort(id TaskID) {
	// todo return error not found

	status := n.task[id]
	if status != PrepareSuccessStatus {
		// todo return error
		return
	}

	n.setTaskStatus(id, AbortStatus)
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
		return ErrTaskFinished
	}

	n.setTaskStatus(id, CommittedStatus)
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

func (n *Node) setTaskStatus(id TaskID, status Status) {
	messages := map[Status]string{
		PrepareSuccessStatus: "prepare %v success",
		PrepareFailedStatus:  "prepare %v failed",
		CommittedStatus:      "commit %v success",
		AbortStatus:          "abort %v success",
	}

	if msg, exist := messages[status]; exist {
		n.addToLog(msg, id)
		n.task[id] = status
	} else {
		panic("unknown status")
	}
}
