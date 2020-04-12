package diameter_of_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func depth(root *TreeNode, d *int) int {
	if root == nil {
		return 0
	}
	l := depth(root.Left, d)
	r := depth(root.Right, d)
	*d = max(*d, l+r+1)
	return max(l, r) + 1
}

func diameterOfBinaryTree(root *TreeNode) int {
	d := 1
	depth(root, &d)
	return d - 1
}
