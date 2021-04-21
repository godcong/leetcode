package _0363

import "math"

/*
363. 矩形区域不超过 K 的最大数值和
给你一个 m x n 的矩阵 matrix 和一个整数 k ，找出并返回矩阵内部矩形区域的不超过 k 的最大数值和。

题目数据保证总会存在一个数值和不超过 k 的矩形区域。



示例 1：


输入：matrix = [[1,0,1],[0,-2,3]], k = 2
输出：2
解释：蓝色边框圈出来的矩形区域 [[0, 1], [-2, 3]] 的数值和是 2，且 2 是不超过 k 的最大数字（k = 2）。
示例 2：

输入：matrix = [[2,2,-1]], k = 3
输出：3


提示：

m == matrix.length
n == matrix[i].length
1 <= m, n <= 100
-100 <= matrix[i][j] <= 100
-105 <= k <= 105


进阶：如果行数远大于列数，该如何设计解决方案？
*/
func maxSumSubmatrix(matrix [][]int, k int) int {
	n := len(matrix)
	m := len(matrix[0])
	max := math.MinInt32
	var sumArry []int
	for i := 0; i < m; i++ {
		sumArry = make([]int, n)
		for j := i; j < m; j++ {
			for o := 0; o < n; o++ {
				sumArry[o] += matrix[o][j]
			}

			max = maxInt(max, subMax(sumArry, k))
		}
	}

	return max
}

func subMax(sumArry []int, k int) int {
	size := len(sumArry)
	max := math.MinInt32
	sum := 0
	for _, v := range sumArry {
		if sum >= 0 {
			sum += v
		} else {
			sum = v
		}
		if max < sum {
			max = sum
		}
	}

	if max <= k {
		return max
	}
	max = math.MinInt32
	for i := range sumArry {
		sum = 0
		for j := i; j < size; j++ {
			sum += sumArry[j]
			if sum <= k && sum > max {
				max = sum
			}
		}
	}

	return max
}

func maxInt(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
