package old

/*
77. 组合
给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。

示例:

输入: n = 4, k = 2
输出:
[
  [2,4],
  [3,4],
  [2,3],
  [1,2],
  [1,3],
  [1,4],
]
*/
func combine(n int, k int) [][]int {
	var temp []int
	var ret [][]int
	var combineDFS func(int)
	combineDFS = func(cur int) {
		if len(temp)+(n-cur+1) < k {
			return
		}
		if len(temp) == k {
			comb := make([]int, k)
			copy(comb, temp)
			ret = append(ret, comb)
			return
		}
		temp = append(temp, cur)
		combineDFS(cur + 1)
		temp = temp[:len(temp)-1]
		combineDFS(cur + 1)
	}
	combineDFS(1)
	return ret
}
