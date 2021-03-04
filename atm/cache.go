package atm

type Coin uint

func (c Coin) ToSum() Sum {
	return Sum(c)
}

type Count uint

type Sum uint

type Cache map[Coin]Count

func (c Cache) Empty() bool {
	if len(c) == 0 {
		return true
	}

	for _, count := range c {
		if count != 0 {
			return false
		}
	}

	return true
}

func (c Cache) Total() Sum {
	var total uint
	for coin, count := range c {
		total += uint(coin) * uint(count)
	}
	return Sum(total)
}
