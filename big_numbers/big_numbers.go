package big_numbers

import (
	"strconv"
)

// todo mayme unit
func NumberFromInt(v int) Number {
	return Number(strconv.Itoa(v))
}

// todo implement
func NumberFromString(v string) (Number, error) {
	panic("not implemented")
}

type Number string

// todo i change to uint
func (n Number) Digit(i int) int {
	// todo check range limit
	// todo check conversation

	if i > n.LastIndex() {
		return 0
	}

	// todo remove !!!
	if i < 0 {
		return 0
	}

	v, _ := strconv.Atoi(string(n[i]))
	return v
}

func (n Number) Len() int {
	return len(n)
}

func (n Number) ToBegin(v int) Number {
	return NumberFromInt(v) + n
}

func (n Number) LastIndex() int {
	return n.Len() - 1
}

func Sum(aug Number, add Number) Number {
	if aug.Len() < add.Len() {
		aug, add = add, aug
	}

	var result Number
	var memo int

	for i1 := aug.LastIndex(); i1 >= 0; i1-- {
		n1 := aug.Digit(i1)

		i2 := i1 - (aug.Len() - add.Len())
		n2 := add.Digit(i2)

		n1 += memo

		s := n1 + n2

		if s > 9 {
			result = result.ToBegin(s - 10)
			memo = 1
			continue
		}

		memo = 0

		result = result.ToBegin(s)

	}

	if memo == 1 {
		result = result.ToBegin(1)
	}

	return result
}
