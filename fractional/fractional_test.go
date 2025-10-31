package fractional_calc

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	t.Run("fractional", func(t *testing.T) {
		f1 := Frac(2, 3)
		f2 := Frac(4, 5)
		f3 := f1.Plus(f2)
		n, d := f3.Parts()
		require.Equal(t, 22, n)
		require.Equal(t, 15, d)
		require.EqualValues(t, 1.4666666666666666, f3.ToFloat64())
	})

	tests := []struct {
		name      string
		a         Fractional
		b         Fractional
		operation string
		result    Fractional
	}{
		{
			name:      "sum zero and number",
			a:         Number(0),
			b:         Number(33),
			operation: "plus",
			result:    Number(33),
		},
		{
			name:      "sum two numbers",
			a:         Number(11),
			b:         Number(33),
			operation: "plus",
			result:    Number(44),
		},
		{
			name:      "sum two fractions",
			a:         Frac(2, 3),
			b:         Frac(4, 5),
			operation: "plus",
			result:    Frac(22, 15),
		},
		{
			name:      "sum fractions and number",
			a:         Number(11),
			b:         Frac(4, 5),
			operation: "plus",
			result:    Frac(59, 5),
		},
		{
			name:      "minus two numbers",
			a:         Number(11),
			b:         Number(22),
			operation: "minus",
			result:    Number(-11),
		},
		{
			name:      "minus two fractions",
			a:         Frac(2, 3),
			b:         Frac(4, 5),
			operation: "minus",
			result:    Frac(-2, 15),
		},
		{
			name:      "sum fractions and number",
			a:         Number(11),
			b:         Frac(4, 5),
			operation: "minus",
			result:    Frac(51, 5),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			switch test.operation {
			case "plus":
				require.Equal(t, test.result, test.a.Plus(test.b))
			case "minus":
				require.Equal(t, test.result, test.a.Minus(test.b))
			default:
				require.Fail(t, "unexpected operation")
			}
		})
	}
}
