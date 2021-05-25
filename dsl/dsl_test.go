package dsl

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDSL(t *testing.T) {
	t.Run("solar without name", func(t *testing.T) {
		require := require.New(t)

		SolarSystem("Sun", func() {
			Description("This is my home world.")
		})

		require.Equal("Sun", Root.Name)
		require.Equal("This is my home world.", Root.Description)
	})
	t.Run("solar with name", func(t *testing.T) {
		require := require.New(t)

		SolarSystem("Sun", func() {
			Name("MySun")
			Description("This is my home world.")
		})

		require.Equal("MySun", Root.Name)
		require.Equal("This is my home world.", Root.Description)
	})
	t.Run("solar with one planet", func(t *testing.T) {
		require := require.New(t)

		SolarSystem("Sun", func() {
			Name("MySun")
			Description("This is my home world.")
			Planet("Earth", func() {
				Description("This my planet.")
				Mass(9999)
			})
		})

		require.Equal("MySun", Root.Name)
		require.Equal("This is my home world.", Root.Description)
		require.Len(Root.Planets, 1)
		require.Contains(Root.Planets, &PlanetNode{
			Name:        "Earth",
			Description: "This my planet.",
			Mass:        9999,
		})
	})
}
