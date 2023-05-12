package celebrity

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPerson(t *testing.T) {
	jane := NewPerson()
	mike := NewPerson()
	bob := NewPerson()

	jane.Show(mike)
	jane.Show(bob)

	require.True(t, jane.IsKnow(mike))
	require.True(t, jane.IsKnow(bob))
	require.False(t, mike.IsKnow(jane))
	require.False(t, bob.IsKnow(jane))

	require.False(t, jane.IsDontKnow(mike))
	require.False(t, jane.IsDontKnow(bob))
	require.True(t, mike.IsDontKnow(jane))
	require.True(t, bob.IsDontKnow(jane))
}
