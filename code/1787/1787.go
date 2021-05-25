package _1787

/*
1787. 使所有区间的异或结果为零
给你一个整数数组 nums​​​ 和一个整数 k​​​​​ 。区间 [left, right]（left <= right）的 异或结果 是对下标位于 left 和 right（包括 left 和 right ）之间所有元素进行 XOR 运算的结果：nums[left] XOR nums[left+1] XOR ... XOR nums[right] 。

返回数组中 要更改的最小元素数 ，以使所有长度为 k 的区间异或结果等于零。



示例 1：

输入：nums = [1,2,0,3,0], k = 1
输出：3
解释：将数组 [1,2,0,3,0] 修改为 [0,0,0,0,0]
示例 2：

输入：nums = [3,4,5,2,1,7,3,4,7], k = 3
输出：3
解释：将数组 [3,4,5,2,1,7,3,4,7] 修改为 [3,4,7,3,4,7,3,4,7]
示例 3：

输入：nums = [1,2,4,1,2,5,1,2,6], k = 3
输出：3
解释：将数组[1,2,4,1,2,5,1,2,6] 修改为 [1,2,3,1,2,3,1,2,3]


提示：

1 <= k <= nums.length <= 2000
0 <= nums[i] < 210
*/
func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func minChanges(nums []int, k int) int {
	fs := make([]map[int]int, k)
	for i := 0; i < k; i++ {
		fs[i] = map[int]int{}
	}
	mas, mass := make([]int, k), make([]int, k)
	for i, n := range nums {
		fs[i%k][n]++
	}

	for i, f := range fs {
		for _, c := range f {
			mas[i] = max(mas[i], c)
		}
	}
	mass[k-1] = mas[k-1]
	for i := k - 2; i >= 0; i-- {
		mass[i] = mass[i+1] + mas[i]
	}

	m := mas[0]
	for i := 1; i < k; i++ {
		m = min(m, mas[i])
	}
	o := mass[0] - m

	var cal func(int, int, int)
	cal = func(i, s, c int) {
		if i == k {
			if s == 0 {
				o = max(o, c)
			}
		} else if c+mass[i] > o {
			for j, f := range fs[i] {
				cal(i+1, s^j, c+f)
			}
		}
	}

	cal(0, 0, 0)
	return len(nums) - o
}