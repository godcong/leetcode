package leetcode

/*
493. 翻转对
给定一个数组 nums ，如果 i < j 且 nums[i] > 2*nums[j] 我们就将 (i, j) 称作一个重要翻转对。

你需要返回给定数组中的重要翻转对的数量。

示例 1:

输入: [1,3,2,3,1]
输出: 2
示例 2:

输入: [2,4,3,5,1]
输出: 3
注意:

给定数组的长度不会超过50000。
输入数组中的所有数字都在32位整数的表示范围内。
*/
func reversePairs(nums []int) int {
	var reversePairsMergeSort func(nums []int, l, r int) int
	reversePairsMergeSort = func(nums []int, l, r int) int {
		if l >= r {
			return 0
		}

		m := (l + r) >> 1

		count := reversePairsMergeSort(nums, l, m) + reversePairsMergeSort(nums, m+1, r)
		cache := make([]int, r-l+1)

		i, k := l, 0

		for j, idx := m+1, l; j <= r; j++ {
			for ; i <= m && nums[i] < nums[j]; i++ {
				cache[k] = nums[i]
				k++
			}

			cache[k] = nums[j]
			k++

			for idx <= m && (nums[idx]+1)>>1 <= nums[j] {
				idx++
			}

			count += m - idx + 1
		}

		copy(nums[l+k:], nums[i:m+1])
		copy(nums[l:], cache[:k])
		return count
	}
	return reversePairsMergeSort(nums, 0, len(nums)-1)
}
