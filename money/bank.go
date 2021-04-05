package money

func NewBank() *Bank {
	return &Bank{}
}

type Bank struct{}

func (b *Bank) Reduce(expr IExpression, currency string) *Money {
	// todo stub
	return NewMoney(24, "USD")
}
