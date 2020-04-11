package diameter_of_binary_tree

import (
	"testing"
)

type test struct {
	input  []int
	expect int
}

func makeTree(arr *[]int, i int) *TreeNode {
	if i >= len(*arr) {
		return nil
	}
	node := TreeNode{}
	node.Val = (*arr)[i]
	node.Left = makeTree(arr, i*2+1)
	node.Right = makeTree(arr, i*2+2)
	return &node
}

func TestDiameterOfBinaryTree(t *testing.T) {
	tests := []test{
		{[]int{1, 2, 3, 4, 5}, 3},
		{[]int{1, 2, 3}, 2},
	}

	for _, c := range tests {
		tree := makeTree(&c.input, 0)
		if result := diameterOfBinaryTree(tree); result != c.expect {
			t.Fatalf("diameterOfBinaryTree(%v) fails.\nExpected %v\nReceived %v", c.input, c.expect, result)
		}
	}
}
