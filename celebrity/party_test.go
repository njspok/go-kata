package celebrity

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParty_Celebrity(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		jane := NewPerson()
		mike := NewPerson()
		bob := NewPerson()
		alice := NewPerson()

		mike.Show(jane)
		bob.Show(jane)
		alice.Show(jane)

		party := NewParty()
		party.Add(jane)
		party.Add(mike)
		party.Add(bob)
		party.Add(alice)

		require.Equal(t, jane, party.Celebrity())
	})
	t.Run("nobody knows anybody", func(t *testing.T) {
		jane := NewPerson()
		mike := NewPerson()
		bob := NewPerson()
		alice := NewPerson()

		party := NewParty()
		party.Add(jane)
		party.Add(mike)
		party.Add(bob)
		party.Add(alice)

		require.Nil(t, party.Celebrity())
	})
	t.Run("cycle", func(t *testing.T) {
		jane := NewPerson()
		mike := NewPerson()
		bob := NewPerson()
		alice := NewPerson()

		jane.Show(mike)
		mike.Show(bob)
		bob.Show(alice)
		alice.Show(jane)

		party := NewParty()
		party.Add(jane)
		party.Add(mike)
		party.Add(bob)
		party.Add(alice)

		require.Nil(t, party.Celebrity())
	})
	t.Run("two pseudo celebrity", func(t *testing.T) {
		jane := NewPerson()
		mike := NewPerson()

		bob := NewPerson()
		alice := NewPerson()

		bob.Show(jane)
		alice.Show(jane)
		bob.Show(mike)
		alice.Show(mike)

		party := NewParty()
		party.Add(jane)
		party.Add(mike)
		party.Add(bob)
		party.Add(alice)

		require.Nil(t, party.Celebrity())
	})
}
