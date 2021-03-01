package libra

func NewLibra() *Libra {
	return &Libra{}
}

type Libra struct {
	left  uint
	right uint
}

func (l *Libra) Balance() uint {
	return diff(l.left, l.right)
}

func (l *Libra) Put(v uint) {
	switch {
	case v == 0:
		return
	case l.isEquilibrium():
		l.putLeft(v)
	case l.rightOverweight():
		l.putLeft(v)
	case l.leftOverweight():
		l.putRight(v)
	default:
		panic("something wrong")
	}
}

func (l *Libra) Left() uint {
	return l.left
}

func (l *Libra) Right() uint {
	return l.right
}

func (l *Libra) putLeft(v uint) {
	l.left += v
}

func (l *Libra) putRight(v uint) {
	l.right += v
}

func (l *Libra) isEquilibrium() bool {
	return l.left == l.right
}

func (l *Libra) rightOverweight() bool {
	return l.left < l.right
}

func (l *Libra) leftOverweight() bool {
	return l.left > l.right
}

func diff(a, b uint) uint {
	if a == b {
		return 0
	}

	if a < b {
		return b - a
	}

	return a - b
}
