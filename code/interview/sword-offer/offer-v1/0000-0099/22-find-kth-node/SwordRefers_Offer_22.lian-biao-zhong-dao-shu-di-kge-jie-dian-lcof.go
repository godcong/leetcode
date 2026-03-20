package _SwordRefers_Offer_22

import (
	. "github.com/godcong/leetcode/common"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getKthFromEnd(head *ListNode, k int) *ListNode {
	n := 0
	var ret *ListNode
	for node := head; node != nil; node = node.Next {
		n++
	}
	for ret = head; n > k; n-- {
		ret = ret.Next
	}
	return ret
}
