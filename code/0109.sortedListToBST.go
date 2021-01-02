package code

/*
109. 有序链表转换二叉搜索树
给定一个单链表，其中的元素按升序排序，将其转换为高度平衡的二叉搜索树。

本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。

示例:

给定的有序链表： [-10, -3, 0, 5, 9],

一个可能的答案是：[0, -3, 9, -10, null, 5], 它可以表示下面这个高度平衡二叉搜索树：

      0
     / \
   -3   9
   /   /
 -10  5
*/
func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	val := []int{head.Val}
	for head.Next != nil {
		head = head.Next
		val = append(val, head.Val)
	}
	root := &TreeNode{}
	sortedListToBSTTreeNode(root, val)
	return root
}

func sortedListToBSTTreeNode(root *TreeNode, val []int) {
	size := len(val)
	if size/2 < size {
		root.Val = val[size/2]
	}
	if len(val[:size/2]) > 0 {
		root.Left = &TreeNode{}
		sortedListToBSTTreeNode(root.Left, val[:size/2])
	}
	if len(val[size/2+1:]) > 0 {
		root.Right = &TreeNode{}
		sortedListToBSTTreeNode(root.Right, val[size/2+1:])
	}
}
