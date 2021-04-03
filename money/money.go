package money

func NewMoney(v uint) *Money {
	return &Money{
		amount: v,
	}
}

type Money struct {
	amount uint
}

func (d *Money) Times(v uint) *Money {
	return NewMoney(d.amount * v)
}

func (d *Money) Amount() uint {
	return d.amount
}

func (d *Money) Equals(o interface{}) bool {
	money := o.(interface{ Amount() uint })
	return d.Amount() == money.Amount()
}
