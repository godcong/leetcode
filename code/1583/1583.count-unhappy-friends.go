package _1583

func unhappyFriends(n int, preferences [][]int, pairs [][]int) int {
	order := make([][]int, n)
	ans := 0
	for i, preference := range preferences {
		order[i] = make([]int, n)
		for j, p := range preference {
			order[i][p] = j
		}
	}
	match := make([]int, n)
	for _, p := range pairs {
		match[p[0]] = p[1]
		match[p[1]] = p[0]
	}

	for x, y := range match {
		index := order[x][y]
		for _, u := range preferences[x][:index] {
			v := match[u]
			if order[u][x] < order[u][v] {
				ans++
				break
			}
		}
	}
	return ans

}
