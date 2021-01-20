package code

import "fmt"

/*
628. 三个数的最大乘积
给定一个整型数组，在数组中找出由三个数组成的最大乘积，并输出这个乘积。

示例 1:

输入: [1,2,3]
输出: 6
示例 2:

输入: [1,2,3,4]
输出: 24
注意:

给定的整型数组长度范围是[3,104]，数组中所有的元素范围是[-1000, 1000]。
输入的数组中任意三个数的乘积不会超出32位有符号整数的范围。
*/
func maximumProduct(nums []int) int {
	min1, min2 := 1>>31, 1>>31
	max1, max2, max3 := -1<<31, -1<<31, -1<<31
	fmt.Println(min1, min2, max1, max2, max3)
	for _, v := range nums {
		if v < min1 {
			min2 = min1
			min1 = v
		} else if v < min2 && v >= min1 {
			min2 = v
		}
		if v > max1 { //是否比最大的数大
			max3 = max2
			max2 = max1
			max1 = v
		} else {          //比最大的数小
			if v > max2 { //是否有第二大的数大
				max3 = max2
				max2 = v
			} else {
				if v > max3 { //是否有第三大的数大
					max3 = v
				}
			}
		}

	}
	if min1*min2*max1 >= max1*max2*max3 {
		return min1 * min2 * max1
	}
	return max1 * max2 * max3

}
