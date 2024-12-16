package addtwonumbers

// 2. Add Two Numbers

// You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit.
// Add the two numbers and return the sum as a linked list.
// You may assume the two numbers do not contain any leading zero, except the number 0 itself.

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var res = &ListNode{}
	var curr = res
	var carry = 0

	// calculate l1 and l2 if not null
	// also this edge case
	// calculate carry also when carry is more than 0
	// for specific case like 7+9 = 16
	for l1 != nil || l2 != nil || carry > 0 {
		// v1 is value of L1 list node, the value is assigned if only l1 is not nil
		v1 := 0
		if l1 != nil {
			v1 = l1.Val
		}

		// v2 is value of L2 list node, the value is assigned if only l2 is not nil
		v2 := 0
		if l2 != nil {
			v2 = l2.Val
		}

		// calculate v1 + v2 and extra carry
		total := v1 + v2 + carry
		// calculate carry with divide by 10
		carry = total / 10
		// calculate value with modulus to get the second value from total
		total = total % 10

		// assign total into linkNode into curr.Next
		curr.Next = &ListNode{Val: total}

		// assign l1 with l1.Next if l1 is exist
		if l1 != nil {
			l1 = l1.Next
		}

		// assign l2 with l2.Next if l2 is exist
		if l2 != nil {
			l2 = l2.Next
		}

		// change current with curr.Next or process next node
		curr = curr.Next
	}

	// return res.Next since we assign the value in res.Next
	return res.Next
}
