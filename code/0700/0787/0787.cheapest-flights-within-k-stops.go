package _0787

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	costBand := 2 * n * 10000
	cost := make([]int, n)
	for i := range cost {
		cost[i] = costBand
	}
	cost[src] = 0

	for i := 0; i <= k; i++ {
		tempCost := append(cost[:0:0], cost...)
		for _, route := range flights {
			tempCost[route[1]] = min(tempCost[route[1]], cost[route[0]]+route[2])
		}
		cost = append(tempCost[:0:0], tempCost...)
	}

	if cost[dst] >= costBand {
		return -1
	}
	return cost[dst]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
