package money

func NewDollar(v uint) Dollar {
	return Dollar{
		amount: v,
	}
}

type Dollar struct {
	amount uint
}

func (d *Dollar) Times(v uint) Dollar {
	return NewDollar(d.amount * v)
}

func (d *Dollar) Amount() uint {
	return d.amount
}
