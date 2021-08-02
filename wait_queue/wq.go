package wait_queue

import "sync"

type Result interface{}

func NewTask(do func() Result) *Task {
	t := &Task{
		do: do,
	}
	t.wg.Add(1)
	return t
}

type Task struct {
	wg     sync.WaitGroup
	do     func() Result
	Result Result
}

func (t *Task) Run() {
	go func() {
		t.Result = t.do()
		t.wg.Done()
	}()
}

func (t *Task) Wait() {
	t.wg.Wait()
}

func NewWaitQueue() *WaitQueue {
	return &WaitQueue{}
}

type WaitQueue struct {
	mu   sync.Mutex
	list sync.Map
}

func (q *WaitQueue) Run(name string, do func() Result) Result {
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