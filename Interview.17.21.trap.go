package leetcode

/*
面试题 17.21. 直方图的水量
给定一个直方图(也称柱状图)，假设有人从上面源源不断地倒水，最后直方图能存多少水量?直方图的宽度为 1。



上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的直方图，在这种情况下，可以接 6 个单位的水（蓝色部分表示水）。 感谢 Marcos 贡献此图。

示例:

输入: [0,1,0,2,1,0,1,3,2,1,2,1]
输出: 6
*/
func trap(height []int) int {
	var ans int
	n := len(height)
	if n == 0 {
		return ans
	}

	leftMax := make([]int, n)
	leftMax[0] = height[0]
	for i := 1; i < n; i++ {
		leftMax[i] = leftMax[i-1]
		if leftMax[i-1] < height[i] {
			leftMax[i] = height[i]
		}
	}

	rightMax := make([]int, n)
	rightMax[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = rightMax[i+1]
		if rightMax[i+1] < height[i] {
			rightMax[i] = height[i]
		}
	}

	for i, h := range height {
		ans += trapMin(leftMax[i], rightMax[i]) - h
	}
	return ans
}

func trapMin(i int, i2 int) int {
	if i < i2 {
		return i
	}
	return i2
}
