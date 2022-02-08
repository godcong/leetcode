package _0391

import (
	"math"
)

func isRectangleCover(rectangles [][]int) bool {
	pos1 := math.MaxInt32
	pos2 := math.MaxInt32
	pos3 := math.MinInt32
	pos4 := math.MinInt32
	area := 0
	set := map[int]bool{}
	for _, rect := range rectangles {
		pos1 = min(pos1, rect[0])
		pos2 = min(pos2, rect[1])
		pos3 = max(pos3, rect[2])
		pos4 = max(pos4, rect[3])
		area += (rect[2] - rect[0]) * (rect[3] - rect[1])
		setPut(set, rect[0], rect[1])
		setPut(set, rect[0], rect[3])
		setPut(set, rect[2], rect[3])
		setPut(set, rect[2], rect[1])
	}
	if (pos3-pos1)*(pos4-pos2) != area {
		return false
	}

	if len(set) != 4 {
		return false
	}

	if set[pos1*10000000+pos2] && set[pos1*10000000+pos4] && set[pos3*10000000+pos4] && set[pos3*10000000+pos2] {
		return true
	}
	return false
}

func setPut(set map[int]bool, x, y int) {
	if set[x*10000000+y] {
		delete(set, x*10000000+y)
	} else {
		set[x*10000000+y] = true
	}
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
