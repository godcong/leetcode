package leetcode

/*
679. 24 点游戏

你有 4 张写有 1 到 9 数字的牌。你需要判断是否能通过 *，/，+，-，(，) 的运算得到 24。

示例 1:

输入: [4, 1, 8, 7]
输出: True
解释: (8-4) * (7-1) = 24

示例 2:

输入: [1, 2, 1, 2]
输出: False

注意:

    除法运算符 / 表示实数除法，而不是整数除法。例如 4 / (1 - 2/3) = 12 。
    每个运算符对两个数进行运算。特别是我们不能用 - 作为一元运算符。例如，[1, 1, 1, 1] 作为输入时，表达式 -1 - 1 - 1 - 1 是不允许的。
    你不能将数字连接在一起。例如，输入为 [1, 2, 1, 2] 时，不能写成 12 + 12 。
*/

const TargetMax = 24 + 1e-6
const TargetMin = 24 - 1e-6

func judgePoint24(nums []int) bool {
	h, i, j, k := 0, 0, 0, 0
	for h = 0; h < 4; h++ {
		for i = 0; i < 4; i++ {
			for j = 0; j < 4; j++ {
				for k = 0; k < 4; k++ {
					if h == i || i == j || j == k || k == h || h == j || i == k {
						continue
					}
					if judgePoint24Cal4(float64(nums[h]), float64(nums[i]), float64(nums[j]), float64(nums[k])) {
						return true
					}
				}
			}
		}
	}
	return false
}

func judgePoint24Cal3(a, b, c float64) bool {
	return judgePoint24Cal(a+b, c) ||
		judgePoint24Cal(a-b, c) ||
		judgePoint24Cal(a*b, c) ||
		judgePoint24Cal(a/b, c) ||
		judgePoint24Cal(b-a, c) ||
		judgePoint24Cal(b/a, c) ||
		judgePoint24Cal(a, b+c) ||
		judgePoint24Cal(a, b-c) ||
		judgePoint24Cal(a, b*c) ||
		judgePoint24Cal(a, b/c) ||
		judgePoint24Cal(a, c-b) ||
		judgePoint24Cal(a, c/b) ||
		judgePoint24Cal(a+c, b) ||
		judgePoint24Cal(a-c, b) ||
		judgePoint24Cal(a*c, b) ||
		judgePoint24Cal(a/c, b) ||
		judgePoint24Cal(c-a, b) ||
		judgePoint24Cal(c/a, b)
}

func judgePoint24Cal4(h, i, j, k float64) bool {
	return judgePoint24Cal3(h+i, j, k) ||
		judgePoint24Cal3(h-i, j, k) ||
		judgePoint24Cal3(h*i, j, k) ||
		judgePoint24Cal3(h/i, j, k) ||
		judgePoint24Cal3(i-h, j, k) ||
		judgePoint24Cal3(i/h, j, k)
}

func judgePoint24Cal(a, b float64) bool {
	switch {
	case a+b < TargetMax && a+b > TargetMin:
	case a*b < TargetMax && a*b > TargetMin:
	case a-b < TargetMax && a-b > TargetMin:
	case b-a < TargetMax && b-a > TargetMin:
	case a/b < TargetMax && a/b > TargetMin:
	case b/a < TargetMax && b/a > TargetMin:
	default:
		return false
	}
	return true
}
