package maximum_depth_of_binary_tree

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

func nodeDepth(node *TreeNode, depth int) int {
	if node.Left == nil && node.Right == nil {
		return depth
	}

	var leftDepth, rightDepth int
	if node.Left != nil {
		leftDepth = nodeDepth(node.Left, depth + 1)
	}
	if node.Right !=nil {
		rightDepth = nodeDepth(node.Right, depth + 1);
	}

	return max(leftDepth, rightDepth)
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return nodeDepth(root, 1)
}
