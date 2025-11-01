package fractional_calc

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	tests := []string{
		"12+0=12",
		"11+22=33",
		"2|3+4|5=22|15",
		"11+4|5=59|5",

		"0-11=-11",
		"11-22=-11",
		"2|3-4|5=-2|15",
		"11-4|5=51|5",
	}
	for _, test := range tests {
		t.Run(test, func(t *testing.T) {
			a, b, c, operation, err := ParseExpression(test)
			require.NoError(t, err)
			switch operation {
			case "+":
				require.Equal(t, c, a.Plus(b))
			case "-":
				require.Equal(t, c, a.Minus(b))
			default:
				require.Fail(t, "unknown operation")
			}
		})
	}
}
