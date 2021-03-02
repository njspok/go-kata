package atm

func NewATM() *ATM {
	return &ATM{
		stat: make(Stat),
	}
}

type Stat map[uint]uint

type ATM struct {
	stat Stat
}

func (a *ATM) Put(coin uint) error {
	a.stat[coin]++
	return nil
}

func (a *ATM) Total() uint {
	var sum uint
	for coin, count := range a.stat {
		sum += coin * count
	}
	return sum
}

func (a *ATM) Stat() Stat {
	return a.stat
}
