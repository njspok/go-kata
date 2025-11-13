package fractional_calc

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCalc(t *testing.T) {
	tests := []string{
		// plus
		"12+0=12",
		"11+22=33",
		"2|3+4|5=22|15",
		"11+4|5=59|5",
		"2|5+1|3=11|15",

		// minus
		"0-11=-11",
		"11-22=-11",
		"2|3-4|5=-2|15",
		"11-4|5=51|5",

		// mult
		"0*5=0",
		"1*1|2=1|2",
		"1|3*1|2=1|6",
		"-4|5*-5|6=20|30",

		// div
		"3|5:4|9=27|20",
	}
	for _, test := range tests {
		t.Run(test, func(t *testing.T) {
			a, b, c, operation, err := ParseCalcExpression(test)
			require.NoError(t, err)
			switch operation {
			case "+":
				require.Equal(t, c, a.Plus(b))
			case "-":
				require.Equal(t, c, a.Minus(b))
			case "*":
				require.Equal(t, c, a.Mult(b))
			case ":":
				require.Equal(t, c, a.Div(b))
			default:
				require.Fail(t, "unknown operation")
			}
		})
	}
}

func TestCmp(t *testing.T) {
	tests := []string{
		"1|3=1|3",
		"1|3=10|30",
	}
	for _, test := range tests {
		t.Run(test, func(t *testing.T) {
			a, b, operation, err := ParseCmpExpression(test)
			require.NoError(t, err)
			switch operation {
			case "=":
				require.True(t, a.Equal(b))
			default:
				require.Fail(t, "unknown operation")
			}
		})
	}
}

func TestFracFromString(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		f, err := FracFromString("11", "22")
		require.NoError(t, err)
		require.Equal(t, Frac(11, 22), f)
	})
	t.Run("invalid string", func(t *testing.T) {
		_, err := FracFromString("invalid string", "99")
		require.Error(t, err)
	})
}
