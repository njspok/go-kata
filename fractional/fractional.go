package fractional_calc

import (
	"regexp"
	"strconv"

	"github.com/pkg/errors"
)

var rx = regexp.MustCompile(`^(\d+)\|?(\d*)([+\-*/])(\d+)\|?(\d*)(=)(-?\d+)\|?(\d*)$`)

func Frac(num, denum int) Fractional {
	return Fractional{
		num:   num,
		denum: denum,
	}
}

func Number(i int) Fractional {
	return Fractional{
		num:   i,
		denum: 1,
	}
}

type Fractional struct {
	num   int
	denum int
}

func (f Fractional) ToFloat64() float64 {
	return float64(f.num) / float64(f.denum)
}

func (f Fractional) Plus(m Fractional) Fractional {
	return Fractional{
		num:   f.num*m.denum + m.num*f.denum,
		denum: f.denum * m.denum,
	}
}

func (f Fractional) Minus(m Fractional) Fractional {
	return Fractional{
		num:   f.num*m.denum - m.num*f.denum,
		denum: f.denum * m.denum,
	}
}

func (f Fractional) Parts() (num, denum int) {
	return f.num, f.denum
}

func FracFromString(num, denum string) (Fractional, error) {
	n, err := strconv.Atoi(num)
	if err != nil {
		return Fractional{}, err
	}

	d := 1
	if denum != "" {
		d, err = strconv.Atoi(denum)
		if err != nil {
			return Fractional{}, err
		}
	}

	return Frac(n, d), nil
}

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
