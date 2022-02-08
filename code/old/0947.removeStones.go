package old

/*
947. 移除最多的同行或同列石头
n 块石头放置在二维平面中的一些整数坐标点上。每个坐标点上最多只能有一块石头。

如果一块石头的 同行或者同列 上有其他石头存在，那么就可以移除这块石头。

给你一个长度为 n 的数组 stones ，其中 stones[i] = [xi, yi] 表示第 i 块石头的位置，返回 可以移除的石子 的最大数量。



示例 1：

输入：stones = [[0,0],[0,1],[1,0],[1,2],[2,1],[2,2]]
输出：5
解释：一种移除 5 块石头的方法如下所示：
1. 移除石头 [2,2] ，因为它和 [2,1] 同行。
2. 移除石头 [2,1] ，因为它和 [0,1] 同列。
3. 移除石头 [1,2] ，因为它和 [1,0] 同行。
4. 移除石头 [1,0] ，因为它和 [0,0] 同列。
5. 移除石头 [0,1] ，因为它和 [0,0] 同行。
石头 [0,0] 不能移除，因为它没有与另一块石头同行/列。
示例 2：

输入：stones = [[0,0],[0,2],[1,1],[2,0],[2,2]]
输出：3
解释：一种移除 3 块石头的方法如下所示：
1. 移除石头 [2,2] ，因为它和 [2,0] 同行。
2. 移除石头 [2,0] ，因为它和 [0,0] 同列。
3. 移除石头 [0,2] ，因为它和 [0,0] 同行。
石头 [0,0] 和 [1,1] 不能移除，因为它们没有与另一块石头同行/列。
示例 3：

输入：stones = [[0,0]]
输出：0
解释：[0,0] 是平面上唯一一块石头，所以不可以移除它。


提示：

1 <= stones.length <= 1000
0 <= xi, yi <= 104
不会有两块石头放在同一个坐标点上
*/
func removeStones(stones [][]int) int {

	arr := make([]int, len(stones))
	for i, _ := range arr {
		arr[i] = i
	}
	rowMp := make(map[int]int)
	colMp := make(map[int]int)
	for k, v := range stones {
		_, ok1 := rowMp[v[0]]
		_, ok2 := colMp[v[1]]
		if ok1 {
			removeStonesUnion(arr, rowMp[v[0]], k)
		} else {
			rowMp[v[0]] = k
		}
		if ok2 {
			removeStonesUnion(arr, colMp[v[1]], k)
		} else {
			colMp[v[1]] = k
		}
	}
	res := 0
	for k, v := range arr {
		if k != v {
			res++
		}
	}
	return res
}

func removeStonesUnion(arr []int, i, j int) int {
	a, b := removeStonesFind(arr, i), removeStonesFind(arr, j)
	if a != b {
		arr[b] = a
		return 1
	}
	return 0
}

func removeStonesFind(arr []int, i int) int {
	for arr[i] != i {
		i = arr[i]
	}
	return i
}
