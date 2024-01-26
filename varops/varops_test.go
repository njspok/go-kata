package varops

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	var result VarValue
	var err error

	o := NewSolver()

	o.Set("Length", func(_ Solver) VarValue { return 10 })
	o.Set("Width", func(_ Solver) VarValue { return 20 })
	o.Set("Height", func(_ Solver) VarValue { return 5 })

	o.Set("FoundationSquare", func(list Solver) VarValue {
		length, err := list.Solve("Length")
		if err != nil {
			panic(err)
		}

		width, err := list.Solve("Width")
		if err != nil {
			panic(err)
		}

		return length * width
	})
	result, err = o.Solve("FoundationSquare")
	require.NoError(t, err)
	require.EqualValues(t, 200, result)

	o.Set("TotalSquare", func(list Solver) VarValue {
		fs, err := list.Solve("FoundationSquare")
		if err != nil {
			panic(err)
		}

		length, err := list.Solve("Length")
		if err != nil {
			panic(err)
		}

		width, err := list.Solve("Width")
		if err != nil {
			panic(err)
		}

		height, err := list.Solve("Height")
		if err != nil {
			panic(err)
		}

		return 2 * (fs + length*height + width*height)
	})
	result, err = o.Solve("TotalSquare")
	require.NoError(t, err)
	require.EqualValues(t, 700, result)

	o.Set("Volume", func(list Solver) VarValue {
		fs, err := list.Solve("FoundationSquare")
		if err != nil {
			panic(err)
		}

		height, err := list.Solve("Height")
		if err != nil {
			panic(err)
		}

		return fs * height
	})
	result, err = o.Solve("Volume")
	require.NoError(t, err)
	require.EqualValues(t, 1000, result)
}
