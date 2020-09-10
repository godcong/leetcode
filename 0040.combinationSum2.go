package leetcode

import "sort"

/*
40. 组合总和 II
给定一个数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的每个数字在每个组合中只能使用一次。

说明：

所有数字（包括目标数）都是正整数。
解集不能包含重复的组合。
示例 1:

输入: candidates = [10,1,2,7,6,1,5], target = 8,
所求解集为:
[
  [1, 7],
  [1, 2, 5],
  [2, 6],
  [1, 1, 6]
]
示例 2:

输入: candidates = [2,5,2,1,2], target = 5,
所求解集为:
[
  [1,2,2],
  [5]
]
*/
func combinationSum2(candidates []int, target int) [][]int {
	var ret [][]int
	sort.Ints(candidates)
	var combinationSum2DFS func([]int, int, int)
	combinationSum2DFS = func(c []int, begin, t int) {
		if t == 0 {
			ret = append(ret, append([]int{}, c...))
			return
		}
		c = append(c, 0)
		for i := begin; i < len(candidates); i++ {
			if i != begin && candidates[i] == candidates[i-1] {
				continue
			}
			if t < candidates[i] {
				return
			}

			c[len(c)-1] = candidates[i]
			combinationSum2DFS(c, i+1, t-candidates[i])
		}
	}

	combinationSum2DFS(nil, 0, target)
	return ret
}
