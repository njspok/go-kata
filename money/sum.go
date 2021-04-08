package money

func NewSum(augend, addend IExpression) *Sum {
	return &Sum{
		augend: augend,
		addend: addend,
	}
}

type Sum struct {
	augend IExpression
	addend IExpression
}

func (s *Sum) Reduce(bank *Bank, to string) *Money {
	total := s.augend.Reduce(bank, to).Amount() + s.addend.Reduce(bank, to).Amount()
	return NewMoney(total, to)
}

func (s *Sum) Plus(addend IExpression) IExpression {
	return NewSum(s, addend)
}

func (s *Sum) Times(multiply uint) IExpression {
	return NewSum(s.augend.Times(multiply), s.addend.Times(multiply))
}
