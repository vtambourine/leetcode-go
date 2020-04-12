package add_two_numbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	node := &ListNode{}
	head := node
	var carry int
	var x, y int

	for l1 != nil || l2 != nil {
		x = 0
		y = 0

		if l1 != nil {
			x = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			y = l2.Val
			l2 = l2.Next
		}

		sum := x + y + carry
		carry = sum / 10

		node.Next = &ListNode{Val: sum % 10}
		node = node.Next
	}

	if carry != 0 {
		node.Next = &ListNode{Val: carry}
		node = node.Next
	}

	return head.Next
}
