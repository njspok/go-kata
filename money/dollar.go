package money

func NewDollar(v uint) *Dollar {
	return &Dollar{
		NewMoney(v, "USD"),
	}
}

type Dollar struct {
	*Money
}
