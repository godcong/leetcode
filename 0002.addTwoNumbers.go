package leetcode

/*
2. 两数相加
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807
*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		l1, l2 = l2, l1
	}
	addTwoNumbersAdd(l1, l2, 0)
	return l1
}

func addTwoNumbersAdd(l1 *ListNode, l2 *ListNode, carry int) {
	if l1 == nil {
		return
	}
	if l2 != nil || carry > 0 {
		carry += l1.Val
		if l2 != nil {
			carry += l2.Val
			l2 = l2.Next
		}
		l1.Val = carry % 10
		carry /= 10
		if l1.Next == nil && l2 != nil {
			l1.Next, l2 = l2, l1.Next
		}
		if l1.Next == nil {
			if l2 != nil {
				l1.Next, l2 = l2, l1.Next
			}
			if l1.Next == nil && carry > 0 {
				l1.Next = new(ListNode)
			}
		}
		addTwoNumbersAdd(l1.Next, l2, carry)
	}
	return
}
