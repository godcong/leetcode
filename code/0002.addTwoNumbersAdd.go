package code

func addTwoNumbersAdd(l1 *ListNode, l2 *ListNode) *ListNode {
	tail := new(ListNode)
	curr := tail
	carry := 0

	for l1 != nil || l2 != nil || carry > 0 {
		curr.Next = new(ListNode)
		curr = curr.Next
		if l1 != nil {
			carry += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			carry += l2.Val
			l2 = l2.Next
		}
		curr.Val = carry % 10
		carry /= 10
	}
	return tail.Next
}
