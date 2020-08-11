package leetcode

//Definition for a TreeNode.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//Definition for a Node.
type Node struct {
	Val       int
	Neighbors []*Node
}
