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
	t.Run("solar with planets", func(t *testing.T) {
		require := require.New(t)

		SolarSystem("Sun", func() {
			Name("MySun")
			Description("This is my home world.")

			Planet("Earth", func() {
				Description("This my home planet.")
				Mass(9999)
			})

			Planet("Mars", func() {
				Description("This my feature planet.")
				Mass(8888)
			})
		})

		require.Equal("MySun", Root.Name)
		require.Equal("This is my home world.", Root.Description)
		require.Len(Root.Planets, 2)
		require.Contains(Root.Planets, &PlanetNode{
			Name:        "Earth",
			Description: "This my home planet.",
			Mass:        9999,
		})
		require.Contains(Root.Planets, &PlanetNode{
			Name:        "Mars",
			Description: "This my feature planet.",
			Mass:        8888,
		})
	})
	t.Run("solar with planets and satellites", func(t *testing.T) {
		require := require.New(t)

		SolarSystem("Sun", func() {
			Name("MySun")
			Description("This is my home world.")

			Planet("Earth", func() {
				Description("This my home planet.")
				Mass(9999)

				Satellite("Moon", func() {
					Description("Beautiful thing!")
					Mass(111)
				})
			})

			Planet("Mars", func() {
				Description("This my feature planet.")
				Mass(8888)

				Satellite("Deimos", func() {
					Description("Rock")
					Mass(222)
				})

				Satellite("Phobos", func() {
					Description("Dead")
					Mass(121)
				})
			})
		})

		require.Equal("MySun", Root.Name)
		require.Equal("This is my home world.", Root.Description)
		require.Len(Root.Planets, 2)
		require.Contains(Root.Planets, &PlanetNode{
			Name:        "Earth",
			Description: "This my home planet.",
			Mass:        9999,
			Satellites: []*SatelliteNode{
				{
					Name:        "Moon",
					Description: "Beautiful thing!",
					Mass:        111,
				},
			},
		})
		require.Contains(Root.Planets, &PlanetNode{
			Name:        "Mars",
			Description: "This my feature planet.",
			Mass:        8888,
			Satellites: []*SatelliteNode{
				{
					Name:        "Deimos",
					Description: "Rock",
					Mass:        222,
				},
				{
					Name:        "Phobos",
					Description: "Dead",
					Mass:        121,
				},
			},
		})
	})
}
