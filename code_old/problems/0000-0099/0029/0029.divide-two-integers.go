package _0029

import (
	"math"
)

func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}

	diffSign := false
	if (dividend < 0) != (divisor < 0) {
		diffSign = true
	}

	i, j, sum := 0, 0, 0
	var temp int
	for {
		temp = 0
		if diffSign {
			temp = sum - divisor<<j
		} else {
			temp = sum + divisor<<j
		}
		if (dividend > 0 && temp > dividend) || (dividend < 0 && temp < dividend) {
			if j == 0 {
				break
			}
			j--
			continue
		}
		sum = temp
		i += 1 << j
		j++
	}

	if diffSign {
		return -i
	}
	return i
}
