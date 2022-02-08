package old

import (
	"math"
)

/*
1584. 连接所有点的最小费用
给你一个points 数组，表示 2D 平面上的一些点，其中 points[i] = [xi, yi] 。

连接点 [xi, yi] 和点 [xj, yj] 的费用为它们之间的 曼哈顿距离 ：|xi - xj| + |yi - yj| ，其中 |val| 表示 val 的绝对值。

请你返回将所有点连接的最小总费用。只有任意两点之间 有且仅有 一条简单路径时，才认为所有点都已连接。



示例 1：



输入：points = [[0,0],[2,2],[3,10],[5,2],[7,0]]
输出：20
解释：

我们可以按照上图所示连接所有点得到最小总费用，总费用为 20 。
注意到任意两个点之间只有唯一一条路径互相到达。
示例 2：

输入：points = [[3,12],[-2,5],[-4,1]]
输出：18
示例 3：

输入：points = [[0,0],[1,1],[1,0],[-1,1]]
输出：4
示例 4：

输入：points = [[-1000000,-1000000],[1000000,1000000]]
输出：4000000
示例 5：

输入：points = [[0,0]]
输出：0


提示：

1 <= points.length <= 1000
-106 <= xi, yi <= 106
所有点 (xi, yi) 两两不同。
*/
func minCostConnectPoints(points [][]int) int {
	visited := make([]bool, len(points))
	minDist := make([]int, len(points))
	for i := 1; i < len(minDist); i++ {
		minDist[i] = math.MaxInt32
	}

	res := 0

	for _ = range visited {
		nextPoint := -1
		d := math.MaxInt32
		for j := range visited {
			if !visited[j] && minDist[j] < d {
				d = minDist[j]
				nextPoint = j
			}
		}

		res += d

		visited[nextPoint] = true

		for j := range visited {
			if !visited[j] {
				minDist[j] = minCostConnectPointsMin(minDist[j], minCostConnectPointsDistance(points[j], points[nextPoint]))
			}
		}
	}

	return res
}

func minCostConnectPointsDistance(p1, p2 []int) int {
	return minCostConnectPointsAbs(p1[0]-p2[0]) + minCostConnectPointsAbs(p1[1]-p2[1])
}

func minCostConnectPointsAbs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func minCostConnectPointsMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
