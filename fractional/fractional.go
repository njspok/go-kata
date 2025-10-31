package fractional_calc

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

func (f Fractional) Parts() (num, denum int) {
	return f.num, f.denum
}
