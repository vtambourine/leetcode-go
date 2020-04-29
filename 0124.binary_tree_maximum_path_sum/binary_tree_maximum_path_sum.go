package binary_tree_maximum_path_sum

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var answer int

func max(n ...int) int {
	m := n[0]
	for _, n := range n {
		if n > m {
			m = n
		}
	}
	return m
}

func dp(root *TreeNode) int {
	left := math.MinInt32
	right := math.MinInt32
	if root.Left != nil {
		left = dp(root.Left)
	}
	if root.Right != nil {
		right = dp(root.Right)
	}
	path := max(root.Val, root.Val+left, root.Val+right)
	answer = max(answer, path, root.Val+left+right)
	return path

}

func maxPathSum(root *TreeNode) int {
	answer = math.MinInt32
	dp(root)
	return answer
}
