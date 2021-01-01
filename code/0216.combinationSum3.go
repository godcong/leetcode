package code

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
	var combinationSum3DFS func([]int, int, int, int)
	combinationSum3DFS = func(c []int, idx, begin, t int) {
		if t == 0 && idx == k {
			ret = append(ret, append([]int{}, c...))
			return
		}
		if idx >= k {
			return
		}
		if c == nil {
			c = make([]int, k, k)
		}
		for i := begin; i <= 9; i++ {
			if i > t {
				return
			}
			c[idx] = i
			combinationSum3DFS(c, idx+1, i+1, t-i)
		}
	}

	combinationSum3DFS(nil, 0, 1, n)
	return ret
}
