package atm

type Coin uint

type Cache map[Coin]uint

func (c Cache) Total() uint {
	var sum uint
	for coin, count := range c {
		sum += uint(coin) * count
	}
	return sum
}
