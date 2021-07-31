package wait_queue

import "sync"

type Result interface{}

func NewWaitQueue() *WaitQueue {
	return &WaitQueue{
		list: make(map[string]*Task),
	}
}

type WaitQueue struct {
	mu   sync.Mutex
	list map[string]*Task
}

func RunTask(do func() Result) *Task {
	task := &Task{}
	task.Add(1)

	go func() {
		task.Result = do()
		task.Done()
	}()

	return task
}

type Task struct {
	sync.WaitGroup

	subs   sync.WaitGroup
	Result Result
}

func (q *WaitQueue) Run(name string, do func() Result) Result {
	q.mu.Lock()

	if task, ok := q.list[name]; ok {
		task.subs.Add(1)
		q.mu.Unlock()

		task.Wait()
		task.subs.Done()

		return task.Result
	}

	task := RunTask(do)
	q.list[name] = task
	q.mu.Unlock()

	task.Wait()

	q.mu.Lock()
	task.subs.Wait()
	delete(q.list, name)
	q.mu.Unlock()

	return task.Result
}
