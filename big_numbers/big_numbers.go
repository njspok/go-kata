package big_numbers

import (
	"strconv"
)

func Sum(aug string, add string) string {
	if len(aug) < len(add) {
		aug, add = add, aug
	}

	var result string
	var memo int

	for i1 := len(aug) - 1; i1 >= 0; i1-- {
		n1, _ := strconv.Atoi(string(aug[i1]))

		i2 := i1 - (len(aug) - len(add))
		n2 := 0
		if i2 >= 0 {
			n2, _ = strconv.Atoi(string(add[i2]))
		}

		n1 += memo

		s := n1 + n2

		if s > 9 {
			result = strconv.Itoa(s-10) + result
			memo = 1
			continue
		}

		memo = 0

		result = strconv.Itoa(s) + result
	}

	if memo == 1 {
		result = "1" + result
	}

	return result
}
