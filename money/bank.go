package money

func NewBank() *Bank {
	return &Bank{}
}

type Bank struct{}

func (b *Bank) Reduce(expr IExpression, currency string) *Money {
	switch expr.(type) {
	case *Sum:
		sum := expr.(*Sum)
		return sum.Reduce(currency)
	case *Money:
		money := expr.(*Money)
		return money.Reduce(currency)
	}

	panic("unknown type expr")
}
