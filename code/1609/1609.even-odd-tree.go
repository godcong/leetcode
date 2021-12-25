package _1609

import (
	"math"
	"runtime/debug"

	. "github.com/godcong/leetcode/common"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isEvenOddTree(root *TreeNode) bool {
	queue := []*TreeNode{root}
	level := 0
	for queue != nil {
		temp := queue
		queue = nil
		prev := 0
		if level%2 == 1 {
			prev = math.MaxInt32
		}
		for _, cur := range temp {
			val := cur.Val
			if val%2 == level%2 || level%2 == 0 && val <= prev || level%2 == 1 && val >= prev {
				return false
			}
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
			prev = val
		}
		level++
	}
	return true
}

func init() { debug.SetGCPercent(-1) }
