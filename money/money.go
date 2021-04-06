package money

func NewMoney(amount uint, currency string) *Money {
	return &Money{
		amount:   amount,
		currency: currency,
	}
}

func NewDollar(v uint) *Money {
	return NewMoney(v, "USD")
}

func NewFrank(v uint) *Money {
	return NewMoney(v, "CHF")
}

type IExpression interface {
	Reduce(currency string) *Money
}

type Money struct {
	amount   uint
	currency string
}

func (m *Money) Times(v uint) *Money {
	return NewMoney(m.amount*v, m.currency)
}

func (m *Money) Amount() uint {
	return m.amount
}

func (m *Money) Currency() string {
	return m.currency
}

func (m *Money) Equals(money *Money) bool {
	return (m.Amount() == money.Amount()) &&
		(m.Currency() == money.Currency())
}

func (m *Money) Plus(money *Money) IExpression {
	return NewSum(m, money)
}

func (m *Money) Reduce(currency string) *Money {
	return NewMoney(m.amount, currency)
}
