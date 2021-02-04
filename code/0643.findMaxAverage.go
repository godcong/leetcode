package code

/*
643. 子数组最大平均数 I
给定 n 个整数，找出平均数最大且长度为 k 的连续子数组，并输出该最大平均数。



示例：

输入：[1,12,-5,-6,50,3], k = 4
输出：12.75
解释：最大平均数 (12-5-6+50)/4 = 51/4 = 12.75


提示：

1 <= k <= n <= 30,000。
所给数据范围 [-10,000，10,000]。
*/
func findMaxAverage(nums []int, k int) float64 {
	tempSum := 0
	var MaxAvg float64 = -10000
	for j := 0; j < k; j++ {
		tempSum += nums[j]
	}
	if v := float64(tempSum) / float64(k); v > MaxAvg {
		MaxAvg = v
	}

	for i := 1; i <= len(nums)-k; i++ {
		tempSum = tempSum - nums[i-1] + nums[i+k-1]
		if v := float64(tempSum) / float64(k); v > MaxAvg {
			MaxAvg = v
		}
	}
	return MaxAvg
}
