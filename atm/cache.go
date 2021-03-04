package atm

type Cache map[uint]uint

func (c Cache) Total() uint {
	var sum uint
	for coin, count := range c {
		sum += coin * count
	}
	return sum
}
