package fractional_calc

import (
	"strconv"
)

// todo check denum is zero !!!
func Frac(num, denum int) Fractional {
	return Fractional{
		num:   num,
		denum: denum,
	}
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

func (f Fractional) Mult(m Fractional) Fractional {
	return Fractional{
		num:   f.num * m.num,
		denum: f.denum * m.denum,
	}
}

func (f Fractional) Div(m Fractional) Fractional {
	return Fractional{
		num:   f.num * m.denum,
		denum: f.denum * m.num,
	}
}

func (f Fractional) Parts() (num, denum int) {
	return f.num, f.denum
}

func (f Fractional) Equal(m Fractional) bool {
	return f.Minus(m).num == 0
}
