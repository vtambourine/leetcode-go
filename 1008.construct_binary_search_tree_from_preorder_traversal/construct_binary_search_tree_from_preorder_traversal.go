package construct_binary_search_tree_from_preorder_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bstFromPreorder(preorder []int) *TreeNode {
	var curr *TreeNode
	head := &TreeNode{preorder[0], nil, nil}

	for i := 1; i < len(preorder); i++ {
		curr = head
		next := TreeNode{preorder[i], nil, nil}

		for curr != nil {
			if next.Val < curr.Val {
				if curr.Left == nil {
					curr.Left = &next
					curr = &next
				}
				curr = curr.Left
			} else {
				if curr.Right == nil {
					curr.Right = &next
					curr = &next
				}
				curr = curr.Right
			}
		}
	}

	return head
}
