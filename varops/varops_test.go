package varops

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	t.Run("calc box parameters", func(t *testing.T) {
		var result VarValue
		var err error

		o := NewSolver()

		// todo
		// set value то что известно, может быть только одно
		// set operation то что вычисляется через другие операции,
		// set operatiion может быть несколько на 1 значение

		o.SetValue("Length", 10)
		o.SetValue("Width", 20)
		o.SetValue("Height", 5)

		o.SetOperation("FoundationSquare", func(list Solver) VarValue {
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

		o.SetOperation("TotalSquare", func(list Solver) VarValue {
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

		o.SetOperation("Volume", func(list Solver) VarValue {
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
	})
}
