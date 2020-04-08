package middle_of_the_linked_list

import (
	"reflect"
	"testing"
)

type test struct {
	input  []int
	expect []int
}

func createList(arr []int) *ListNode {
	var prevNode *ListNode
	head := ListNode{}
	currNode := &head
	for _, v := range arr {
		currNode.Val = v
		if prevNode != nil {
			prevNode.Next = currNode
		}
		prevNode = currNode
		currNode = &ListNode{}
	}
	return &head
}

func serializeList(head *ListNode) []int {
	result := make([]int, 0)
	node := head
	for node != nil {
		result = append(result, node.Val)
		node = node.Next
	}
	return result
}

func TestMiddleNode(t *testing.T) {
	tests := []test{
		{[]int{1, 2, 3, 4, 5}, []int{3, 4, 5}},
		{[]int{1, 2, 3, 4, 5, 6}, []int{4, 5, 6}},
	}

	for _, c := range tests {
		result := serializeList(middleNode(createList(c.input)))
		if !reflect.DeepEqual(result, c.expect) {
			t.Fatalf("middleNode((%v) fails.\nExpected %v\nReceived %v", c.input, c.expect, result)
		}
	}
}
