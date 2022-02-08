package _1001

func gridIllumination(n int, lamps [][]int, queries [][]int) []int {
	m := len(lamps)
	type pair struct{ x, y int }
	points := make(map[pair]struct{}, m)
	row := make(map[int]int, m)
	col := make(map[int]int, m)
	diagonal := make(map[int]int, m)
	antiDiagonal := make(map[int]int, m)
	for _, lamp := range lamps {
		x, y := lamp[0], lamp[1]
		p := pair{x, y}
		if _, ok := points[p]; ok {
			continue
		}
		points[p] = struct{}{}
		row[x]++
		col[y]++
		diagonal[x-y]++
		antiDiagonal[x+y]++
	}

	ans := make([]int, len(queries))
	for i, query := range queries {
		r, c := query[0], query[1]
		if row[r] > 0 || col[c] > 0 || diagonal[r-c] > 0 || antiDiagonal[r+c] > 0 {
			ans[i] = 1
		}
		for x := r - 1; x <= r+1; x++ {
			for y := c - 1; y <= c+1; y++ {
				if x < 0 || y < 0 || x >= n || y >= n {
					continue
				}
				if _, ok := points[pair{x, y}]; !ok {
					continue
				}
				delete(points, pair{x, y})
				row[x]--
				col[y]--
				diagonal[x-y]--
				antiDiagonal[x+y]--
			}
		}
	}
	return ans
}
