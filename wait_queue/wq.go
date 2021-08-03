package wait_queue

import "sync"

type Result interface{}

type Action func() Result

func NewWaitQueue() *WaitQueue {
	return &WaitQueue{}
}

type WaitQueue struct {
	mu   sync.Mutex
	list sync.Map
}

func (q *WaitQueue) Run(name string, do Action) Result {
	newTask := NewTask(do)
	v, loaded := q.list.LoadOrStore(name, newTask)
	task := v.(*Task)

	if loaded {
		task.Wait()
		return task.Result
	}

	task.Run()
	task.Wait()
	q.list.Delete(name)
	return task.Result
}

func (q *WaitQueue) Load(name string) *Task {
	v, _ := q.list.Load(name)
	if v == nil {
		return nil
	}

	return v.(*Task)
}
