package _0815

import (
	"runtime/debug"
)

func init() { debug.SetGCPercent(-1) }

func numBusesToDestination(routes [][]int, source int, target int) int {
	if source == target {
		return 0
	}

	// 起点的 bus 和 终点的 bus 集合
	var start, end []int
	for busNum, v1 := range routes {
		for _, v2 := range v1 {
			if v2 == source {
				start = append(start, busNum)
			}
			if v2 == target {
				end = append(end, busNum)
			}
		}
	}
	graph := buildGraph(routes)
	min := -1
	for _, s := range start {
		for _, e := range end {
			res := bfs(graph, s, e)
			if min == -1 {
				min = res
			} else if res != -1 && res < min {
				min = res
			}
		}
	}
	if min == -1 {
		return min
	} else {
		return min + 1
	}
}

// 构建邻接表
func buildGraph(routes [][]int) [][]int {
	busCount := len(routes)
	graph := make([][]int, busCount)
	for i := 0; i < busCount; i++ {
		for j := i; j < busCount; j++ {
			if i != j && edge(routes[i], routes[j]) {
				graph[i] = append(graph[i], j)
				graph[j] = append(graph[j], i)
			}
		}
	}
	return graph
}

func edge(busA []int, busB []int) bool {
	i, j := 0, 0
	for i < len(busA) && j < len(busB) {
		if busA[i] == busB[j] {
			return true
		} else if busA[i] < busB[j] {
			i++
		} else {
			j++
		}
	}
	return false
}

// 广度优先搜索找出最短路径
func bfs(graph [][]int, s, e int) int {
	// 如果起始公交和结束公交是同一路，则换乘 0 趟
	if s == e {
		return 0
	}

	length := len(graph)

	// 到达每个节点的路径
	weight := make([]int, length)
	weight[s] = -1

	// 节点队列
	queue := make([]int, 1)
	queue[0] = s

	for i := 0; i < length; i++ {
		if len(queue) <= i {
			return -1
		}
		curr := queue[i]

		// 开始遍历
		for j := 0; j < len(graph[curr]); j++ {
			newVertex := graph[curr][j]
			if weight[newVertex] == 0 {
				queue = append(queue, newVertex)
				weight[newVertex] = i + 1

				if newVertex == e {
					return weight[newVertex]
				}
			}
		}

	}
	return -1
}
