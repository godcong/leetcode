package old

/*
416. 分割等和子集
给定一个只包含正整数的非空数组。是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。

注意:

每个数组中的元素不会超过 100
数组的大小不会超过 200
示例 1:

输入: [1, 5, 11, 5]

输出: true

解释: 数组可以分割成 [1, 5, 5] 和 [11].


示例 2:

输入: [1, 2, 3, 5]

输出: false

解释: 数组不能分割成两个元素和相等的子集.
*/
func canPartition(nums []int) bool {
	totalNum := 0
	for index := range nums {
		totalNum += nums[index]
	}
	if totalNum%2 != 0 {
		return false
	}
	numCount := len(nums)

	row := numCount
	col := totalNum / 2
	dp := make([]bool, col+1)
	dp[0] = true
	for i := 1; i <= row; i++ {
		for w := col; w > 0; w-- {
			if w-nums[i-1] >= 0 {
				dp[w] = dp[w] || dp[w-nums[i-1]]
			}
		}
	}
	return dp[col]
}
