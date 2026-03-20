package _0997

func findJudge(n int, trust [][]int) int {
	d := make([]int, n+1)
	max := 1
	for i := range trust {
		d[trust[i][1]]++
		d[trust[i][0]]--
		if d[trust[i][1]] > d[max] {
			max = trust[i][1]
		}
	}

	if d[max] == n-1 {
		return max
	}
	return -1
}
