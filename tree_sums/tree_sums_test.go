package tree_sums

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNode_Sums(t *testing.T) {
	t.Run("root node", func(t *testing.T) {
		root := &Node{
			Val:   100,
			Left:  nil,
			Right: nil,
		}

		require.Equal(t, 100, root.Sum())

		sum, path := root.MaxPathSum()
		require.Equal(t, 100, sum)
		require.Equal(t, []int{100}, path)
	})
	t.Run("tree", func(t *testing.T) {
		root := &Node{
			Val: 100,
			Left: &Node{
				Val: 99,
				Left: &Node{
					Val:   11,
					Left:  nil,
					Right: nil,
				},
				Right: &Node{
					Val:   14,
					Left:  nil,
					Right: nil,
				},
			},
			Right: &Node{
				Val:   9,
				Left:  nil,
				Right: nil,
			},
		}

		require.Equal(t, 233, root.Sum())

		sum, path := root.MaxPathSum()
		require.Equal(t, 213, sum)
		require.Equal(t, []int{100, 99, 14}, path)
	})
}
