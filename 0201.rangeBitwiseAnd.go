package leetcode

/*
201. 数字范围按位与
给定范围 [m, n]，其中 0 <= m <= n <= 2147483647，返回此范围内所有数字的按位与（包含 m, n 两端点）。

示例 1:

输入: [5,7]
输出: 4
示例 2:

输入: [0,1]
输出: 0
*/
func rangeBitwiseAnd(m int, n int) int {
	//Brian Kernighan 算法
	for m < n {
		n &= n - 1
	}
	return n
}
