package code

/*
834. 树中距离之和
给定一个无向、连通的树。树中有 N 个标记为 0...N-1 的节点以及 N-1 条边 。

第 i 条边连接节点 edges[i][0] 和 edges[i][1] 。

返回一个表示节点 i 与其他所有节点距离之和的列表 ans。

示例 1:

输入: N = 6, edges = [[0,1],[0,2],[2,3],[2,4],[2,5]]
输出: [8,12,6,10,10,10]
解释:
如下为给定的树的示意图：
  0
 / \
1   2
   /|\
  3 4 5

我们可以计算出 dist(0,1) + dist(0,2) + dist(0,3) + dist(0,4) + dist(0,5)
也就是 1 + 1 + 2 + 2 + 2 = 8。 因此，answer[0] = 8，以此类推。
说明: 1 <= N <= 10000
*/
func sumOfDistancesInTree(N int, edges [][]int) []int {
	type vertex struct {
		id, treePathSum, treeSize int
	}
	adjEdges := make([][]int, N)
	for _, e := range edges {
		u, v := e[0], e[1]
		adjEdges[u] = append(adjEdges[u], v)
		adjEdges[v] = append(adjEdges[v], u)
	}

	bfsQueue, parents := make([]*vertex, N), make([]*vertex, N)
	bfsQueue[N-1] = &vertex{0, 0, 1}
	head, tail := N-1, N-1
	for head >= 0 {
		v := bfsQueue[head]
		head--
		for _, u := range adjEdges[v.id] {
			if nil != parents[u] || 0 == u {
				continue
			}

			parents[u] = v
			tail--
			if tail >= 0 {
				bfsQueue[tail] = &vertex{u, 0, 1}
			}
		}
	}

	for i := 0; i < len(bfsQueue)-1; i++ {
		v := bfsQueue[i]
		p := parents[v.id]
		p.treePathSum = p.treePathSum + v.treePathSum + v.treeSize
		p.treeSize = p.treeSize + v.treeSize
	}

	for i := len(bfsQueue) - 2; i >= 0; i-- {
		v := bfsQueue[i]
		p := parents[v.id]
		v.treePathSum = v.treePathSum +
			(p.treePathSum - (v.treePathSum + v.treeSize)) +
			(p.treeSize - v.treeSize)
		v.treeSize = p.treeSize
	}

	result := make([]int, N)
	for _, v := range bfsQueue {
		result[v.id] = v.treePathSum
	}

	return result
}
