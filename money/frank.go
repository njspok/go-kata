package money

func NewFrank(v uint) *Frank {
	return &Frank{
		NewMoney(v),
	}
}

type Frank struct {
	*Money
}
