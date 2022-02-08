package _0797

func allPathsSourceTarget(graph [][]int) [][]int {
	stk := []int{0}
	var ans [][]int
	var dfs func(int)
	dfs = func(x int) {
		if x == len(graph)-1 {
			ans = append(ans, append([]int(nil), stk...))
			return
		}
		for _, y := range graph[x] {
			stk = append(stk, y)
			dfs(y)
			stk = stk[:len(stk)-1]
		}
	}
	dfs(0)
	return ans
}
