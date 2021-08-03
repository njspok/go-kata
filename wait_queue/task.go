package wait_queue

import "sync"

func NewTask(do Action) *Task {
	t := &Task{
		do: do,
	}
	t.wg.Add(1)
	return t
}

type Task struct {
	wg     sync.WaitGroup
	do     Action
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
