package construct_binary_search_tree_from_preorder_traversal

import (
	"reflect"
	"testing"
)

type test struct {
	input  []int
	expect []int
}

func serializeTree(node *TreeNode) []int {
	result := make([]int, 0)
	queue := []*TreeNode{node}
	for len(queue) != 0 {
		front := queue[0]
		queue = queue[1:]
		result = append(result, front.Val)
		if front.Left != nil {
			queue = append(queue, front.Left)
		}
		if front.Right != nil {
			queue = append(queue, front.Right)
		}
	}
	return result
}

func TestBSTFromPreorder(t *testing.T) {
	tests := []test{
		{[]int{8, 5, 1, 7, 10, 12}, []int{8, 5, 10, 1, 7, 12}},
	}

	for _, c := range tests {
		result := serializeTree(bstFromPreorder(c.input))
		if !reflect.DeepEqual(result, c.expect) {
			t.Fatalf("solution(%v) fails.\nExpected %v\nReceived %v", c.input, c.expect, result)
		}
	}
}
