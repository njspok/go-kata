package money

func NewDollar(v uint) *Money {
	return NewMoney(v, "USD")
}
