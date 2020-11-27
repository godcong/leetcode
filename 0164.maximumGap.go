package leetcode

import "math"

/*
164. 最大间距
给定一个无序的数组，找出数组在排序之后，相邻元素之间最大的差值。

如果数组元素个数小于 2，则返回 0。

示例 1:

输入: [3,6,9,1]
输出: 3
解释: 排序后的数组是 [1,3,6,9], 其中相邻元素 (3,6) 和 (6,9) 之间都存在最大差值 3。
示例 2:

输入: [10]
输出: 0
解释: 数组元素个数小于 2，因此返回 0。
说明:

你可以假设数组中所有元素都是非负整数，且数值在 32 位有符号整数范围内。
请尝试在线性时间复杂度和空间复杂度的条件下解决此问题。
*/
func maximumGap(nums []int) int {
	type pair struct{ min, max int }

	ret := 0
	n := len(nums)
	if n < 2 {
		return ret
	}

	max := 0
	min := math.MaxUint32
	for i := range nums {
		if nums[i] > max {
			max = nums[i]
		}
		if nums[i] < min {
			min = nums[i]
		}
	}

	d := (max - min) / (n - 1)
	if d < 1 {
		d = 1
	}

	bucketSize := (max-min)/d + 1

	// 存储 (桶内最小值，桶内最大值) 对，(-1, -1) 表示该桶是空的
	buckets := make([]pair, bucketSize)
	for i := range buckets {
		buckets[i] = pair{-1, -1}
	}
	for _, v := range nums {
		bid := (v - min) / d
		if buckets[bid].min == -1 {
			buckets[bid].min = v
			buckets[bid].max = v
		} else {
			if v < buckets[bid].min {
				buckets[bid].min = v
			}
			if v > buckets[bid].max {
				buckets[bid].max = v
			}
		}
	}

	prev := -1
	for i, b := range buckets {
		if b.min == -1 {
			continue
		}
		if prev != -1 {
			if b.min-buckets[prev].max > ret {
				ret = b.min - buckets[prev].max
			}
		}
		prev = i
	}

	return ret
}
