package code

/*
349. 两个数组的交集
给定两个数组，编写一个函数来计算它们的交集。



示例 1：

输入：nums1 = [1,2,2,1], nums2 = [2,2]
输出：[2]
示例 2：

输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
输出：[9,4]
*/
func intersection(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums1) {
		nums1, nums2 = nums2, nums1
	}

	setNums1 := make(map[int]bool)
	for i := range nums1 {
		setNums1[nums1[i]] = true
	}

	var ret []int
	for i := range nums2 {
		if add, b := setNums1[nums2[i]]; b && add {
			ret = append(ret, nums2[i])
			setNums1[nums2[i]] = false
		}
	}
	return ret
}
