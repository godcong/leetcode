package leetcode

/*
404. 左叶子之和
计算给定二叉树的所有左叶子之和。

示例：

    3
   / \
  9  20
    /  \
   15   7

在这个二叉树中，有两个左叶子，分别是 9 和 15，所以返回 24
*/
func sumOfLeftLeaves(root *TreeNode) int {
	var ret int
	var sumOfLeftLeaves func(*TreeNode)
	sumOfLeftLeaves = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Left != nil && node.Left.Left == nil && node.Left.Right == nil {
			ret += node.Left.Val
		}

		sumOfLeftLeaves(node.Left)
		sumOfLeftLeaves(node.Right)
	}
	sumOfLeftLeaves(root)
	return ret
}
