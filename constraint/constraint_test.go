package constraint

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type State string

type Color string

func NewMapColorConstraint(p1 State, p2 State) *MapColoringConstraint {
	return &MapColoringConstraint{
		place1: p1,
		place2: p2,
	}
}

type MapColoringConstraint struct {
	place1 State
	place2 State
}

func (c *MapColoringConstraint) Satisfied(solution Solution[State, Color]) bool {
	color1, exist1 := solution[c.place1]
	color2, exist2 := solution[c.place2]

	if !exist1 || !exist2 {
		return true
	}

	return color1 != color2
}

func (c *MapColoringConstraint) Variables() []State {
	return []State{c.place1, c.place2}
}

type LinearSystem struct{}

func (l *LinearSystem) Satisfied(solution Solution[string, int]) bool {
	x, e1 := solution["X"]
	y, e2 := solution["Y"]

	if !e1 || !e2 {
		return true
	}

	return ((x + y) == 12) && ((x - y) == -2)
}

func (l *LinearSystem) Variables() []string {
	return []string{"X", "Y"}
}

func TestCSP(t *testing.T) {
	t.Run("coloring map", func(t *testing.T) {
		variables := []State{
			"Western Australia",
			"Northern Territory",
			"South Australia",
			"Queensland",
			"New South Wales",
			"Victoria",
			"Tasmania",
		}

		constraints := []Constraint[State, Color]{
			NewMapColorConstraint("Western Australia", "Northern Territory"),
			NewMapColorConstraint("Western Australia", "South Australia"),
			NewMapColorConstraint("South Australia", "Northern Territory"),
			NewMapColorConstraint("Queensland", "Northern Territory"),
			NewMapColorConstraint("Queensland", "South Australia"),
			NewMapColorConstraint("Queensland", "New South Wales"),
			NewMapColorConstraint("New South Wales", "South Australia"),
			NewMapColorConstraint("Victoria", "South Australia"),
			NewMapColorConstraint("Victoria", "New South Wales"),
			NewMapColorConstraint("Victoria", "Tasmania"),
		}

		t.Run("solved", func(t *testing.T) {
			domains := make(map[State][]Color)
			for _, v := range variables {
				domains[v] = []Color{"red", "green", "blue"}
			}

			csp, err := NewCSP(variables, domains)
			require.NoError(t, err)

			err = csp.AddConstraints(constraints...)
			require.NoError(t, err)

			result := csp.Search()
			require.Equal(t, Solution[State, Color]{
				"Western Australia":  "red",
				"Northern Territory": "green",
				"South Australia":    "blue",
				"Queensland":         "red",
				"New South Wales":    "green",
				"Victoria":           "red",
				"Tasmania":           "green",
			}, result)
		})
		t.Run("not solved", func(t *testing.T) {
			domains := make(map[State][]Color)
			for _, v := range variables {
				domains[v] = []Color{"red", "green"}
			}

			csp, err := NewCSP(variables, domains)
			require.NoError(t, err)

			err = csp.AddConstraints(constraints...)
			require.NoError(t, err)

			result := csp.Search()
			require.Nil(t, result)
		})
	})
	t.Run("linear", func(t *testing.T) {
		csp, err := NewCSP([]string{"X", "Y"}, map[string][]int{
			"X": {-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
			"Y": {-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
		})
		require.NoError(t, err)

		require.NoError(t, csp.AddConstraint(&LinearSystem{}))

		solution := csp.Search()
		require.EqualValues(t, 1, solution)
	})
}
