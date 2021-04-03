package money

func NewDollar(v uint) *Dollar {
	return &Dollar{
		NewMoney(v),
	}
}

type Dollar struct {
	*Money
}
