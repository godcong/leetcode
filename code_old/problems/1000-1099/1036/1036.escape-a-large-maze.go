package _1036

import (
	"sort"
)

type pair struct{ x, y int }

var dirs = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func isEscapePossible(blocked [][]int, source []int, target []int) bool {
	n := len(blocked)
	if n < 2 {
		return true
	}
	rows := []int{source[0], target[0]}
	cols := []int{source[1], target[1]}
	for _, b := range blocked {
		rows = append(rows, b[0])
		cols = append(cols, b[1])
	}

	rMapping, rBound := discrete(rows)
	cMapping, cBound := discrete(cols)

	grid := make([][]bool, rBound+1)
	for i := range grid {
		grid[i] = make([]bool, cBound+1)
	}
	for _, b := range blocked {
		grid[rMapping[b[0]]][cMapping[b[1]]] = true
	}

	sx, sy := rMapping[source[0]], cMapping[source[1]]
	tx, ty := rMapping[target[0]], cMapping[target[1]]
	grid[sx][sy] = true
	q := []pair{{sx, sy}}
	for len(q) > 0 {
		p := q[0]
		q = q[1:]
		for _, d := range dirs {
			x, y := p.x+d.x, p.y+d.y
			if 0 <= x && x <= rBound && 0 <= y && y <= cBound && !grid[x][y] {
				if x == tx && y == ty {
					return true
				}
				grid[x][y] = true
				q = append(q, pair{x, y})
			}
		}
	}
	return false
}

func discrete(blocks []int) (map[int]int, int) {
	sort.Ints(blocks)

	id := 0
	if blocks[0] > 0 {
		id = 1
	}
	mapping := map[int]int{blocks[0]: id}
	pre := blocks[0]
	for _, v := range blocks[1:] {
		if v != pre {
			if v == pre+1 {
				id++
			} else {
				id += 2
			}
			mapping[v] = id
			pre = v
		}
	}

	const boundary int = 1e6
	if blocks[len(blocks)-1] != boundary-1 {
		id++
	}

	return mapping, id
}
