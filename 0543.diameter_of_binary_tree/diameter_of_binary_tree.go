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

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(height(root.Left), height(root.Right))
}

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(
		height(root.Left)+height(root.Right),
		max(diameterOfBinaryTree(root.Left), diameterOfBinaryTree(root.Right)),
	)
}
