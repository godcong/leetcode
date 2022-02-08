package _0677

import (
	"strings"
)

type MapSum struct {
	m map[string]int
}

func Constructor() MapSum {
	return MapSum{m: make(map[string]int)}
}

func (this *MapSum) Insert(key string, val int) {
	this.m[key] = val
}

func (this *MapSum) Sum(prefix string) int {
	sum := 0
	for s, v := range this.m {
		if strings.HasPrefix(s, prefix) {
			sum += v
		}
	}
	return sum
}

/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,m);
 * param_2 := obj.Sum(prefix);
 */
