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
	Reduce(bank *Bank, to string) *Money
	Plus(IExpression) IExpression
	Times(uint) IExpression
}

type Money struct {
	amount   uint
	currency string
}

func (m *Money) Times(v uint) IExpression {
	return NewMoney(m.amount*v, m.currency)
}

func (m *Money) Amount() uint {
	return m.amount
}

func (m *Money) Currency() string {
	return m.currency
}

func (m *Money) Equals(money interface{}) bool {
	m2, ok := money.(*Money)
	if !ok {
		panic("cant compare objects")
	}

	return (m.Amount() == m2.Amount()) &&
		(m.Currency() == m2.Currency())
}

func (m *Money) Plus(addend IExpression) IExpression {
	return NewSum(m, addend)
}

func (m *Money) Reduce(bank *Bank, to string) *Money {
	rate := bank.Rate(m.currency, to)
	return NewMoney(m.Amount()/rate, to)
}
