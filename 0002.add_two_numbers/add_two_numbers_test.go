package add_two_numbers

import (
	"testing"
)

type test struct {
	input  [2][]int
	expect []int
}

func NewList(arr []int) *ListNode {
	var node *ListNode
	for i := len(arr) - 1; i >= 0; i-- {
		node = &ListNode{arr[i], node}
	}
	return node
}

func (list *ListNode) isEqual(anotherList *ListNode) bool {
	leftNode := list
	rightNode := anotherList
	for leftNode != nil && rightNode != nil {
		if leftNode.Val != rightNode.Val {
			return false
		}
		leftNode = leftNode.Next
		rightNode = rightNode.Next
	}
	return leftNode == nil && rightNode == nil
}

func (list *ListNode) serialize() []int {
	node := list
	result := make([]int, 0)
	for node != nil {
		result = append(result, node.Val)
		node = node.Next
	}
	return result
}

func TestAddTwoNumbers(t *testing.T) {
	tests := []test{
		{[2][]int{{2, 4, 3}, {5, 6, 4}}, []int{7, 0, 8}},
		{[2][]int{{1, 1, 1}, {2, 2, 2}}, []int{3, 3, 3}},
		{[2][]int{{1, 1, 1}, {2}}, []int{3, 1, 1}},
	}

	for _, c := range tests {
		list1 := NewList(c.input[0])
		list2 := NewList(c.input[1])
		expect := NewList(c.expect)
		if result := addTwoNumbers(list1, list2); !result.isEqual(expect) {
			t.Fatalf("addTwoNumbers(%v) fails.\nExpected %v\nReceived %v", c.input, c.expect, result.serialize())
		}
	}
}
