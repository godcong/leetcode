package leetcode

/*
327. 区间和的个数
给定一个整数数组 nums，返回区间和在 [lower, upper] 之间的个数，包含 lower 和 upper。
区间和 S(i, j) 表示在 nums 中，位置从 i 到 j 的元素之和，包含 i 和 j (i ≤ j)。

说明:
最直观的算法复杂度是 O(n2) ，请在此基础上优化你的算法。

示例:

输入: nums = [-2,5,-1], lower = -2, upper = 2,
输出: 3
解释: 3个区间分别是: [0,0], [2,2], [0,2]，它们表示的和分别为: -2, -1, 2。
*/
func countRangeSum(nums []int, lower int, upper int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	tmp := make([]int, len(nums)+1)

	s := make([]int, len(nums)+1)
	for i := 1; i <= len(nums); i++ {
		s[i] = s[i-1] + nums[i-1]
	}

	return mergesort(s, 0, len(nums), tmp, lower, upper)
}

func mergesort(s []int, left, right int, tmp []int, lower, upper int) int {
	if left >= right {
		return 0
	}
	res := 0
	mid := (right + left) / 2
	res += mergesort(s, left, mid, tmp, lower, upper)
	res += mergesort(s, mid+1, right, tmp, lower, upper)

	l := left
	low, upp := mid+1, mid+1
	for l <= mid {

		for low <= right && s[low]-s[l] < lower {
			low++
		}

		for upp <= right && s[upp]-s[l] <= upper {
			upp++
		}

		res += upp - low

		l++
	}

	merge(s, tmp, left, mid, right)
	return res
}

func merge(s []int, tmp []int, left, mid, right int) {
	l := left
	r := mid + 1
	index := left
	for l <= mid && r <= right {
		if s[l] <= s[r] {
			tmp[index] = s[l]
			l++
			index++
		} else {
			tmp[index] = s[r]
			r++
			index++
		}
	}
	for l <= mid {
		tmp[index] = s[l]
		l++
		index++
	}
	for r <= right {
		tmp[index] = s[r]
		r++
		index++
	}
	for i := left; i <= right; i++ {
		s[i] = tmp[i]
	}
}
