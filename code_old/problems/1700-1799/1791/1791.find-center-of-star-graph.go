package _1791

func findCenter(edges [][]int) int {
	n := len(edges) + 1
	degrees := make([]int, n+1)
	for _, e := range edges {
		degrees[e[0]]++
		degrees[e[1]]++
	}
	for i, d := range degrees {
		if d == n-1 {
			return i
		}
	}
	return -1
}
