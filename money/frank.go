package money

func NewFrank(v uint) *Frank {
	return &Frank{
		NewMoney(v, "CHF"),
	}
}

type Frank struct {
	*Money
}
