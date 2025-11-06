package fractional_calc

import (
	"regexp"

	"github.com/pkg/errors"
)

// regex for pars strings operations in two fractions like
// 2+2=4 or 2|5+1|3=11|15
var rx = regexp.MustCompile(`^(-?\d+)\|?(\d*)([+\-*:])(-?\d+)\|?(\d*)(=)(-?\d+)\|?(\d*)$`)

func ParseExpression(str string) (a, b, c Fractional, operation string, err error) {
	groups := rx.FindAllStringSubmatch(str, -1)

	if len(groups) != 1 || len(groups[0]) != 9 {
		return a, b, c, operation, errors.New("invalid expression")
	}

	a, err = FracFromString(groups[0][1], groups[0][2])
	if err != nil {
		return a, b, c, operation, err
	}

	operation = groups[0][3]

	b, err = FracFromString(groups[0][4], groups[0][5])
	if err != nil {
		return a, b, c, operation, err
	}

	c, err = FracFromString(groups[0][7], groups[0][8])
	if err != nil {
		return a, b, c, operation, err
	}

	return a, b, c, operation, nil
}
