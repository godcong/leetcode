package _0220

import "math"

/*
220. 存在重复元素 III
给你一个整数数组 nums 和两个整数 k 和 t 。请你判断是否存在 两个不同下标 i 和 j，使得 abs(nums[i] - nums[j]) <= t ，同时又满足 abs(i - j) <= k 。

如果存在则返回 true，不存在返回 false。



示例 1：

输入：nums = [1,2,3,1], k = 3, t = 0
输出：true
示例 2：

输入：nums = [1,0,1,1], k = 1, t = 2
输出：true
示例 3：

输入：nums = [1,5,9,1,5,9], k = 2, t = 3
输出：false


提示：

0 <= nums.length <= 2 * 104
-231 <= nums[i] <= 231 - 1
0 <= k <= 104
0 <= t <= 231 - 1
*/
func getBucketId(idx, width int) int {
	if idx < 0 {
		return (idx+1)/width - 1
	}
	return idx / width
}
func judgeAbsLess(i, j, k int) bool {
	return math.Abs(float64(i)-float64(j)) <= float64(k)
}
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	l := len(nums)
	bucket := make(map[int]int)
	for i := 0; i < l; i++ {
		no := getBucketId(nums[i], t+1)
		if idx, ok := bucket[no]; ok && judgeAbsLess(i, idx, k) {
			return true
		}
		if prev, ok := bucket[no-1]; ok {
			if judgeAbsLess(prev, i, k) && judgeAbsLess(nums[prev], nums[i], t) {
				return true
			}
		}
		if next, ok := bucket[no+1]; ok {
			if judgeAbsLess(next, i, k) && judgeAbsLess(nums[next], nums[i], t) {
				return true
			}
		}
		bucket[no] = i
	}
	return false
}
