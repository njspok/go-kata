package tree_sums

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func (n *Node) MaxPathSum() (int, []int) {
	if n == nil {
		return 0, []int{}
	}

	lsum, lpath := n.Left.MaxPathSum()
	rsum, rpath := n.Right.MaxPathSum()

	path := []int{n.Val}
	sum := n.Val

	if lsum <= rsum {
		sum += rsum
		path = append(path, rpath...)
	} else {
		sum += lsum
		path = append(path, lpath...)
	}

	return sum, path
}

func (n *Node) Sum() int {
	if n == nil {
		return 0
	}

	return n.Val + n.Left.Sum() + n.Right.Sum()
}
