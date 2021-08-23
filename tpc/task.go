package tpc

type TaskID int

func NewTask(id TaskID) *Task {
	return &Task{
		ID: id,
	}
}

type Task struct {
	ID TaskID
}
