package tpc

type TaskID int

func NewTask(id TaskID) *Task {
	return &Task{
		id: id,
	}
}

type Task struct {
	id TaskID
}

func (t *Task) ID() TaskID {
	return t.id
}
