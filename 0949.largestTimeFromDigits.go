package leetcode

import "fmt"

/*
949. 给定数字能组成的最大时间
给定一个由 4 位数字组成的数组，返回可以设置的符合 24 小时制的最大时间。

最小的 24 小时制时间是 00:00，而最大的是 23:59。从 00:00 （午夜）开始算起，过得越久，时间越大。

以长度为 5 的字符串返回答案。如果不能确定有效时间，则返回空字符串。



示例 1：

输入：[1,2,3,4]
输出："23:41"
示例 2：

输入：[5,5,5,5]
输出：""


提示：

A.length == 4
0 <= A[i] <= 9
*/
var largestTimeFromDigitsHourMax = 24
var largestTimeFromDigitsMinuteMax = 60

func largestTimeFromDigits(A []int) string {
	max := [2]int{-1, -1}
	var maxDigit = func(h, m int) {
		if h >= largestTimeFromDigitsHourMax || m >= largestTimeFromDigitsMinuteMax {
			return
		}
		if h > (max[0]) || (h == (max[0]) && m > (max[1])) {
			max[0] = h
			max[1] = m
		}
	}
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			for k := 0; k < 4; k++ {
				for l := 0; l < 4; l++ {
					if i == j || i == k || i == l || j == k || j == l || k == l {
						continue
					}
					maxDigit(A[i]*10+A[j], A[k]*10+A[l])
				}
			}
		}
	}
	if max[0] == -1 || max[1] == -1 {
		return ""
	}
	return fmt.Sprintf("%02v:%02v", max[0], max[1])
}
