package _1719

import (
	"fmt"
	"sort"
)

func checkWays(pairs [][]int) int {
	var adj [501][501]bool
	c := make([]int, 501)
	list := make([]int, 0, 500)
	for _, pair := range pairs {
		u, v := pair[0], pair[1]
		adj[u][v] = true
		adj[v][u] = true
		c[u]++
		c[v]++
		if c[u] == 1 {
			list = append(list, u)
		}
		if c[v] == 1 {
			list = append(list, v)
		}
	}
	n := len(list)
	sort.Slice(list, func(i, j int) bool { return c[list[i]] > c[list[j]] })
	if c[list[0]] < n-1 {
		return 0
	}

	f := make([]int, 501)
	f[list[0]] = -1

	cnt := 0
	for i := 1; i < n; i++ {
		u := list[i]
		for j := i - 1; f[u] == 0 && j >= 0; j-- {
			v := list[j]
			if adj[u][v] {
				f[u] = v
			}
		}
	}
	ret := 1
	for i := 1; i < n; i++ {
		u := list[i]
		v := f[u]
		for v != -1 {
			if !adj[u][v] {
				fmt.Println(u, v)
				return 0
			}
			if c[u] == c[v] {
				ret = 2
			}
			cnt++
			v = f[v]
		}
	}
	if cnt < len(pairs) {
		return 0
	}
	return ret
}
