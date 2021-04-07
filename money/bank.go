package money

import "fmt"

func NewBank() *Bank {
	return &Bank{
		rates: make(map[string]uint),
	}
}

type Bank struct {
	rates map[string]uint
}

func (b *Bank) Reduce(expr IExpression, currency string) *Money {
	return expr.Reduce(b, currency)
}

func (b *Bank) AddRate(from string, to string, i uint) {
	b.rates[b.key(from, to)] = i
}

func (b *Bank) Rate(from, to string) uint {
	if from == to {
		return 1
	}

	if rate, ok := b.rates[b.key(from, to)]; ok {
		return rate
	}

	panic("unknown rates")
}

func (b *Bank) key(from string, to string) string {
	return fmt.Sprintf("%v.%v", from, to)
}
