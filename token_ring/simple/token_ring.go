package simple

// Есть пачка процессов и они объеденины в кольцо.
// Процессы передают друг другу маркер.
// Если процесс принял маркер, и ему он не нужен, то он передает его дальше.
// Если процесс принял маркер и он ему нужен, то он выполняет действие и передает маркер дальше.

func NewRing() *Ring {
	r := &Ring{
		ch: make(chan struct{}, 1),
	}
	r.ch <- struct{}{}
	return r
}

type Ring struct {
	ch chan struct{}
}

func (r *Ring) TakeToken() {
	<-r.ch
}

func (r *Ring) PutToken() {
	r.ch <- struct{}{}
}
