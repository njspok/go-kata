package money

func NewSum(augend, addend *Money) *Sum {
	return &Sum{
		augend: augend,
		addend: addend,
	}
}

type Sum struct {
	augend *Money
	addend *Money
}

func (s *Sum) Reduce(currency string) *Money {
	total := s.augend.Amount() + s.addend.Amount()
	return NewMoney(total, currency)
}
