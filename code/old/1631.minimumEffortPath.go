package old

/*
1631. 最小体力消耗路径
你准备参加一场远足活动。给你一个二维 rows x columns 的地图 heights ，其中 heights[row][col] 表示格子 (row, col) 的高度。一开始你在最左上角的格子 (0, 0) ，且你希望去最右下角的格子 (rows-1, columns-1) （注意下标从 0 开始编号）。你每次可以往 上，下，左，右 四个方向之一移动，你想要找到耗费 体力 最小的一条路径。

一条路径耗费的 体力值 是路径上相邻格子之间 高度差绝对值 的 最大值 决定的。

请你返回从左上角走到右下角的最小 体力消耗值 。



示例 1：



输入：heights = [[1,2,2],[3,8,2],[5,3,5]]
输出：2
解释：路径 [1,3,5,3,5] 连续格子的差值绝对值最大为 2 。
这条路径比路径 [1,2,2,2,5] 更优，因为另一条路径差值最大值为 3 。
示例 2：



输入：heights = [[1,2,3],[3,8,4],[5,3,5]]
输出：1
解释：路径 [1,2,3,4,5] 的相邻格子差值绝对值最大为 1 ，比路径 [1,3,5,3,5] 更优。
示例 3：


输入：heights = [[1,2,1,1,1],[1,2,1,2,1],[1,2,1,2,1],[1,2,1,2,1],[1,1,1,2,1]]
输出：0
解释：上图所示路径不需要消耗任何体力。


提示：

rows == heights.length
columns == heights[i].length
1 <= rows, columns <= 100
1 <= heights[i][j] <= 106
*/
func minimumEffortPath(heights [][]int) int {
	row := len(heights)
	col := len(heights[0])
	left, right := 0, 1000000

	isCheck := make([][]int, row)
	for i := 0; i < row; i++ {
		isCheck[i] = make([]int, col)
	}

	idx := 1
	for left < right {
		mid := (right-left)/2 + left
		if minimumEffortPathDFS(0, 0, heights, mid, isCheck, idx) {
			right = mid
		} else {
			left = mid + 1
		}
		idx++
	}
	return left
}

var minimumEffortPathX = []int{0, 1, 0, -1}
var minimumEffortPathY = []int{1, 0, -1, 0}

func minimumEffortPathDFS(si, sj int, heights [][]int, mid int, checked [][]int, idx int) bool {
	row := len(heights)
	col := len(heights[0])
	if si == row-1 && sj == col-1 {
		return true
	}

	checked[si][sj] = idx
	for i := 0; i < 4; i++ {
		newX := si + minimumEffortPathX[i]
		newY := sj + minimumEffortPathY[i]
		if newX < 0 || newX >= row || newY < 0 || newY >= col || checked[newX][newY] == idx {
			continue
		}
		if minimumEffortPathABS(heights[newX][newY]-heights[si][sj]) > mid {
			continue
		}

		if minimumEffortPathDFS(newX, newY, heights, mid, checked, idx) {
			return true
		}

	}
	return false
}

func minimumEffortPathABS(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
