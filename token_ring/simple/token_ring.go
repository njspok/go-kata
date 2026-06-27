package simple

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
