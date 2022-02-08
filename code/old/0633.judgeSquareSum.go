package old

import "math"

/*
633. 平方数之和
给定一个非负整数 c ，你要判断是否存在两个整数 a 和 b，使得 a^2 + b^2 = c。

示例1:

输入: 5
输出: True
解释: 1 * 1 + 2 * 2 = 5


示例2:

输入: 3
输出: False
*/
func judgeSquareSum(c int) bool {
	var j float64
	for i := 0; i*i <= c; i++ {
		j = math.Sqrt(float64(c - (i * i)))
		if j == float64(int(j)) {
			return true
		}
	}
	return false
}
