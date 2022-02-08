package _0752

import (
	"strconv"
)

func openLock(deadends []string, target string) int {
	if target == "0000" {
		return 0
	}
	var deadArray [10000]bool
	for _, s := range deadends {
		i, _ := strconv.Atoi(s)
		deadArray[i] = true
	}
	targetInt, _ := strconv.Atoi(target)
	if deadArray[0] {
		return -1
	}
	q := []int{0}
	visit := [10000]bool{0: true}
	var tq []int
	helper := func(out []int, x, rate int) {
		if k := (x / rate) % 10; k == 0 {
			out[0] = x + 9*rate
			out[1] = x + 1*rate
		} else if k == 9 {
			out[0] = x - 9*rate
			out[1] = x - 1*rate
		} else {
			out[0] = x + 1*rate
			out[1] = x - 1*rate
		}
	}
	res := 0
	for len(q) > 0 {
		res++
		for len(q) > 0 {
			n := q[len(q)-1]
			q = q[:len(q)-1]

			var next [8]int
			helper(next[:], n, 1)
			helper(next[2:], n, 10)
			helper(next[4:], n, 100)
			helper(next[6:], n, 1000)
			for _, v := range next {
				if v == targetInt {
					return res
				}
				if !deadArray[v] && !visit[v] {
					tq = append(tq, v)
					visit[v] = true
				}
			}
		}
		q, tq = tq, q
		tq = tq[:0]
	}
	return -1
}
