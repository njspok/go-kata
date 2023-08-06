package wait_queue

import "sync"

type Action[T any] func() T

func NewWaitQueue[T any]() *WaitQueue[T] {
	return &WaitQueue[T]{}
}

type WaitQueue[T any] struct {
	mu   sync.Mutex
	list sync.Map
}

func (q *WaitQueue[T]) Run(name string, do Action[T]) T {
	newTask := NewTask(do)
	v, loaded := q.list.LoadOrStore(name, newTask)
	task := v.(*Task[T])

	if loaded {
		task.Wait()
		return task.Result
	}

	task.Run()
	task.Wait()
	q.list.Delete(name)
	return task.Result
}

func (q *WaitQueue[T]) Load(name string) *Task[T] {
	v, _ := q.list.Load(name)
	if v == nil {
		return nil
	}

	return v.(*Task[T])
}
