package _0773

import (
	"strconv"
)

const Target = "123450"

var Neighbor = [6][]int{
	{1, 3}, {0, 2, 4}, {1, 5}, {0, 4}, {1, 3, 5}, {2, 4},
}

func slidingPuzzle(board [][]int) int {
	m := 2
	n := 3

	var begin string

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			v := strconv.Itoa(board[i][j])
			begin += v
		}
	}

	visited := make(map[string]bool)
	visited[begin] = true
	queue := []string{begin}

	step := 0

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			pool := queue[0]
			queue = queue[1:]
			if pool == Target {
				return step
			}

			index0 := 0
			for pool[index0] != '0' {
				index0++
			}
			for _, c := range Neighbor[index0] {
				s := swap(c, index0, pool)
				if !visited[s] {
					visited[s] = true
					queue = append(queue, s)
				}
			}
		}
		step++
	}
	return -1

}

func swap(i, j int, s string) string {
	b := []byte(s)
	b[i], b[j] = b[j], b[i]
	return string(b)
}
