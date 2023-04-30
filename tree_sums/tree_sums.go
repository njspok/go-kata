package tree_sums

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func (n *Node) MaxPathSum() int {
	if n == nil {
		return 0
	}

	return n.Val + max(n.Left.MaxPathSum(), n.Right.MaxPathSum())
}

func (n *Node) Sum() int {
	if n == nil {
		return 0
	}

	return n.Val + n.Left.Sum() + n.Right.Sum()
}

func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}
