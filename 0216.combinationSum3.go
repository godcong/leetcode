package leetcode

/*
216. 组合总和 III
找出所有相加之和为 n 的 k 个数的组合。组合中只允许含有 1 - 9 的正整数，并且每种组合中不存在重复的数字。

说明：

所有数字都是正整数。
解集不能包含重复的组合。
示例 1:

输入: k = 3, n = 7
输出: [[1,2,4]]
示例 2:

输入: k = 3, n = 9
输出: [[1,2,6], [1,3,5], [2,3,4]]
*/
func combinationSum3(k int, n int) [][]int {
	var ret [][]int
	var combinationSum3DFS func([]int, int, int)
	combinationSum3DFS = func(c []int, begin, t int) {
		if t == 0 && len(c) == k {
			ret = append(ret, append([]int{}, c...))
			return
		}
		if len(c) >= k {
			return
		}
		c = append(c, 0)
		for i := begin; i <= 9; i++ {
			if i > t {
				return
			}
			c[len(c)-1] = i
			combinationSum3DFS(c, i+1, t-i)
		}
	}

	combinationSum3DFS(nil, 1, n)
	return ret

}
