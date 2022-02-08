package old

/*
39. 组合总和
给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的数字可以无限制重复被选取。

说明：

所有数字（包括 target）都是正整数。
解集不能包含重复的组合。
示例 1：

输入：candidates = [2,3,6,7], target = 7,
所求解集为：
[
  [7],
  [2,2,3]
]
示例 2：

输入：candidates = [2,3,5], target = 8,
所求解集为：
[
  [2,2,2,2],
  [2,3,3],
  [3,5]
]


提示：

1 <= candidates.length <= 30
1 <= candidates[i] <= 200
candidate 中的每个元素都是独一无二的。
1 <= target <= 500
*/
func combinationSum(candidates []int, target int) [][]int {
	var ret [][]int
	var combinationSumDFS func([]int, int, int)
	combinationSumDFS = func(c []int, begin, i int) {
		if i == 0 {
			ret = append(ret, append([]int{}, c...))
			return
		}
		c = append(c, 0)
		for ; begin < len(candidates); begin++ {
			if i < candidates[begin] {
				continue
			}
			c[len(c)-1] = candidates[begin]
			combinationSumDFS(c, begin, i-candidates[begin])
		}
	}

	combinationSumDFS(nil, 0, target)
	return ret
}
