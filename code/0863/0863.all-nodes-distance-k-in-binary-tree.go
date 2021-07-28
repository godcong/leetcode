package _0863

import (
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
func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	var ans []int
	parents := map[int]*TreeNode{}
	var findParents func(*TreeNode)
	findParents = func(node *TreeNode) {
		if node.Left != nil {
			parents[node.Left.Val] = node
			findParents(node.Left)
		}
		if node.Right != nil {
			parents[node.Right.Val] = node
			findParents(node.Right)
		}
	}
	findParents(root)

	var findAns func(*TreeNode, *TreeNode, int)
	findAns = func(node, from *TreeNode, depth int) {
		if node == nil {
			return
		}
		if depth == k { // 将所有深度为 k 的结点的值计入结果
			ans = append(ans, node.Val)
			return
		}
		if node.Left != from {
			findAns(node.Left, node, depth+1)
		}
		if node.Right != from {
			findAns(node.Right, node, depth+1)
		}
		if parents[node.Val] != from {
			findAns(parents[node.Val], node, depth+1)
		}
	}
	findAns(target, nil, 0)
	return ans
}
