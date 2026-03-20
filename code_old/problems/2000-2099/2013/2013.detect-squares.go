package _2013

import (
	"runtime/debug"
)

const (
	N = 1001
)

type DetectSquares struct {
	m map[int]int
	l [N][]int
}

func Constructor() DetectSquares {
	return DetectSquares{
		m: make(map[int]int, N),
		l: [N][]int{},
	}
}

func (this *DetectSquares) Add(point []int) {
	f := ref(point[0], point[1])
	this.m[f]++
	if this.m[f] == 1 {
		this.l[point[0]] = append(this.l[point[0]], point[1])
	}
}

func (this *DetectSquares) Count(point []int) int {
	ret := 0
	for _, y := range this.l[point[0]] {
		if y == point[1] {
			continue
		}
		ll := abs(y - point[1])
		bas := this.m[ref(point[0], y)]
		if point[0]+ll >= 0 && point[0]+ll < N {
			ret += bas * this.m[ref(point[0]+ll, point[1])] * this.m[ref(point[0]+ll, y)]
		}
		if point[0]-ll >= 0 && point[0]-ll < N {
			ret += bas * this.m[ref(point[0]-ll, point[1])] * this.m[ref(point[0]-ll, y)]
		}
	}
	return ret
}

func ref(x, y int) int {
	return x*N + y
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func init() { debug.SetGCPercent(-1) }

/**
 * Your DetectSquares object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(point);
 * param_2 := obj.Count(point);
 */
