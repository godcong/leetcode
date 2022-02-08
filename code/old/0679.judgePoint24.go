package old

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

const TargetMax = 24 + 1e-3
const TargetMin = 24 - 1e-3

var calList = [6]byte{'+', '-', '*', '/', '_', '?'}

func judgePoint24(nums []int) bool {
	h, i, j, k := 0, 0, 0, 0
	for h = 0; h < 4; h++ {
		for i = 0; i < 4; i++ {
			for j = 0; j < 4; j++ {
				for k = 0; k < 4; k++ {
					if h == i || i == j || j == k || k == h || h == j || i == k {
						continue
					}
					if judgePoint24Cal4(float32(nums[h]), float32(nums[i]), float32(nums[j]), float32(nums[k])) {
						return true
					}
				}
			}
		}
	}
	return false
}

func judgePoint24Cal3(a, b, c float32) bool {
	for idx := range calList {
		if judgePoint24Cal2(judgePoint24Cal(a, b, calList[idx]), c) ||
			judgePoint24Cal2(judgePoint24Cal(a, c, calList[idx]), b) ||
			judgePoint24Cal2(judgePoint24Cal(b, c, calList[idx]), a) {
			return true
		}
	}
	return false
}

func judgePoint24Cal4(h, i, j, k float32) bool {
	for idx := range calList {
		if judgePoint24Cal3(judgePoint24Cal(h, i, calList[idx]), j, k) {
			return true
		}
	}
	return false
}

func judgePoint24Cal2(a, b float32) bool {
	for i := range calList {
		if v := judgePoint24Cal(a, b, calList[i]); v < TargetMax && v > TargetMin {
			return true
		}
	}
	return false
}

func judgePoint24Cal(a, b float32, symb byte) float32 {
	switch symb {
	case '+':
		return a + b
	case '-':
		return a - b
	case '_':
		return b - a
	case '*':
		return a * b
	case '/':
		return a / b
	case '?':
		return b / a
	}
	return 0
}
