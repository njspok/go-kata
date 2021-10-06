package big_numbers

import (
	"fmt"
	"strconv"
)

// todo mayme unit
func NumberFromUint(v uint) Number {
	return Number(fmt.Sprintf("%d", v))
}

// todo implement
func NumberFromString(v string) (Number, error) {
	panic("not implemented")
}

// todo change to []uint8?
type Number string

// todo i change to uint8?
func (n Number) Digit(i uint) uint {
	// todo check range limit
	// todo check conversation

	if i > n.LastIndex() {
		return 0
	}

	// todo remove !!!
	if i < 0 {
		return 0
	}

	v, _ := strconv.Atoi(string(n[n.LastIndex()-i]))
	return uint(v)
}

func (n Number) Len() uint {
	return uint(len(n))
}

func (n Number) ToBegin(v uint) Number {
	return NumberFromUint(v) + n
}

func (n Number) LastIndex() uint {
	return n.Len() - 1
}

func Sum(aug Number, add Number) Number {
	if aug.Len() < add.Len() {
		aug, add = add, aug
	}

	var result Number
	var memo uint

	for i := uint(0); i <= aug.LastIndex(); i++ {
		d1 := aug.Digit(i)
		d2 := add.Digit(i)

		d := d1 + d2 + memo

		if d > 9 {
			d -= 10
			memo = 1
		} else {
			memo = 0
		}

		result = result.ToBegin(d)

	}

	if memo == 1 {
		result = result.ToBegin(1)
	}

	return result
}
