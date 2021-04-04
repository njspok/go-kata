package money

func NewMoney(amount uint, currency string) *Money {
	return &Money{
		amount:   amount,
		currency: currency,
	}
}

type Money struct {
	amount   uint
	currency string
}

func (d *Money) Times(v uint) *Money {
	return NewMoney(d.amount*v, d.currency)
}

func (d *Money) Amount() uint {
	return d.amount
}

func (d *Money) Currency() string {
	return d.currency
}

func (d *Money) Equals(o interface{}) bool {
	money := o.(interface {
		Amount() uint
		Currency() string
	})
	return (d.Amount() == money.Amount()) &&
		(d.Currency() == money.Currency())
}
