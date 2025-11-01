package fractional_calc

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// todo rewrite tests in style "2+3=5", "2/5+5/3=31/15"

func Test2(t *testing.T) {
	tests := []struct {
		name string
		expr string
	}{
		{
			name: "sum with zero",
			expr: "12+0=12",
		},
		{
			name: "sum two numbers",
			expr: "11+22=33",
		},
		{
			name: "sum two fractions",
			expr: "2|3+4|5=22|15",
		},
		{
			name: "sum fraction and number",
			expr: "11+4|5=59|5",
		},
		{
			name: "minus two numbers",
			expr: "11-22=-11",
		},
		{
			name: "minus two fractions",
			expr: "2|3-4|5=-2|15",
		},
		{
			name: "sum fractions and number",
			expr: "11-4|5=51|5",
		},
	}
	for _, test := range tests {
		t.Run(test.name+": "+test.expr, func(t *testing.T) {
			a, b, c, operation, err := ParseExpression(test.expr)
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
