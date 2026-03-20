package _0382

import (
	"math/rand"

	. "github.com/godcong/leetcode/common"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type Solution struct {
	s []int
}

func Constructor(head *ListNode) Solution {
	s := Solution{}
	for p := head; p != nil; p = p.Next {
		s.s = append(s.s, p.Val)
	}
	return s
}

func (this *Solution) GetRandom() int {
	return this.s[rand.Intn(len(this.s))]
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */
