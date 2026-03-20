package _0295

import (
	"runtime/debug"
)

func init() { debug.SetGCPercent(-1) }

type MedianFinder struct {
	queMin, queMax []int
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{
		queMin: make([]int, 0),
		queMax: make([]int, 0),
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if len(this.queMin) == 0 && len(this.queMax) == 0 {
		return -1
	}
	if len(this.queMin) == len(this.queMax) {
		return float64(this.queMin[0]+this.queMax[0]) / 2.0
	} else if len(this.queMin) < len(this.queMax) {
		return float64(this.queMax[0])
	} else {
		return float64(this.queMin[0])
	}
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */

func (this *MedianFinder) AddNum(num int) {
	// 第一个元素，扔到小顶堆
	if len(this.queMin) == 0 {
		this.queMin = append(this.queMin, num)
		return
	}

	if len(this.queMin) < len(this.queMax) {
		// 添加到小顶堆
		if num < this.queMax[0] {
			this.queMin = append(this.queMin, this.queMax[0])
			this.queMinUp()
			this.queMax[0] = num
			this.queMaxDown()
		} else {
			this.queMin = append(this.queMin, num)
			this.queMinUp()
		}
	} else {
		// 添加到大顶堆
		if num > this.queMin[0] {
			this.queMax = append(this.queMax, this.queMin[0])
			this.queMaxUp()
			this.queMin[0] = num
			this.queMinDown()
		} else {
			this.queMax = append(this.queMax, num)
			this.queMaxUp()
		}
	}
}

func (this *MedianFinder) queMinUp() {
	child := len(this.queMin) - 1
	parent := (len(this.queMin) - 1) / 2
	for child > 0 && this.queMin[parent] > this.queMin[child] {
		this.queMin[child], this.queMin[parent] = this.queMin[parent], this.queMin[child]
		child = parent
		parent = (parent - 1) / 2
	}
}

func (this *MedianFinder) queMinDown() {
	parent := 0
	childL := 2*parent + 1
	childR := childL + 1

	for childL < len(this.queMin) {
		min := childL
		if childR < len(this.queMin) && this.queMin[childL] > this.queMin[childR] {
			min = childR
		}
		if this.queMin[min] > this.queMin[parent] {
			min = parent
		}
		if min == parent {
			break
		}
		this.queMin[min], this.queMin[parent] = this.queMin[parent], this.queMin[min]
		parent = min
		childL = 2*parent + 1
		childR = childL + 1
	}
}

func (this *MedianFinder) queMaxUp() {
	child := len(this.queMax) - 1
	parent := (len(this.queMax) - 1) / 2
	for child > 0 && this.queMax[parent] < this.queMax[child] {
		this.queMax[child], this.queMax[parent] = this.queMax[parent], this.queMax[child]
		child = parent
		parent = (parent - 1) / 2
	}
}

func (this *MedianFinder) queMaxDown() {
	parent := 0
	childL := 2*parent + 1
	childR := childL + 1

	for childL < len(this.queMax) {
		min := childL
		if childR < len(this.queMax) && this.queMax[childL] < this.queMax[childR] {
			min = childR
		}
		if this.queMax[min] < this.queMax[parent] {
			min = parent
		}
		if min == parent {
			break
		}
		this.queMax[min], this.queMax[parent] = this.queMax[parent], this.queMax[min]
		parent = min
		childL = 2*parent + 1
		childR = childL + 1
	}
}
