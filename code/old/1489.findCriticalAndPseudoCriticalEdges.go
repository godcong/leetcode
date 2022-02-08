package old

import "sort"

/*
1489. 找到最小生成树里的关键边和伪关键边
给你一个 n 个点的带权无向连通图，节点编号为 0 到 n-1 ，同时还有一个数组 edges ，其中 edges[i] = [fromi, toi, weighti] 表示在 fromi 和 toi 节点之间有一条带权无向边。最小生成树 (MST) 是给定图中边的一个子集，它连接了所有节点且没有环，而且这些边的权值和最小。

请你找到给定图中最小生成树的所有关键边和伪关键边。如果从图中删去某条边，会导致最小生成树的权值和增加，那么我们就说它是一条关键边。伪关键边则是可能会出现在某些最小生成树中但不会出现在所有最小生成树中的边。

请注意，你可以分别以任意顺序返回关键边的下标和伪关键边的下标。



示例 1：



输入：n = 5, edges = [[0,1,1],[1,2,1],[2,3,2],[0,3,2],[0,4,3],[3,4,3],[1,4,6]]
输出：[[0,1],[2,3,4,5]]
解释：上图描述了给定图。
下图是所有的最小生成树。

注意到第 0 条边和第 1 条边出现在了所有最小生成树中，所以它们是关键边，我们将这两个下标作为输出的第一个列表。
边 2，3，4 和 5 是所有 MST 的剩余边，所以它们是伪关键边。我们将它们作为输出的第二个列表。
示例 2 ：



输入：n = 4, edges = [[0,1,1],[1,2,1],[2,3,1],[0,3,1]]
输出：[[],[0,1,2,3]]
解释：可以观察到 4 条边都有相同的权值，任选它们中的 3 条可以形成一棵 MST 。所以 4 条边都是伪关键边。


提示：

2 <= n <= 100
1 <= edges.length <= findCriticalAndPseudoCriticalEdgesUnionMin(200, n * (n - 1) / 2)
edges[i].length == 3
0 <= fromi < toi < n
1 <= weighti <= 1000
所有 (fromi, toi) 数对都是互不相同的。
*/
type findCriticalAndPseudoCriticalEdgesUnion struct {
	parent, size []int
}

func newFindCriticalAndPseudoCriticalEdgesUnion(n int) *findCriticalAndPseudoCriticalEdgesUnion {
	parent := make([]int, n)
	size := make([]int, n)
	for i := range parent {
		parent[i] = i
		size[i] = 1
	}
	return &findCriticalAndPseudoCriticalEdgesUnion{parent, size}
}

func (uf *findCriticalAndPseudoCriticalEdgesUnion) find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *findCriticalAndPseudoCriticalEdgesUnion) union(x, y int) bool {
	fx, fy := uf.find(x), uf.find(y)
	if fx == fy {
		return false
	}
	if uf.size[fx] < uf.size[fy] {
		fx, fy = fy, fx
	}
	uf.size[fx] += uf.size[fy]
	uf.parent[fy] = fx
	return true
}

func findCriticalAndPseudoCriticalEdges(n int, edges [][]int) [][]int {
	m := len(edges)
	edgeType := make([]int, m)

	for i, e := range edges {
		edges[i] = append(e, i)
	}
	sort.Slice(edges, func(i, j int) bool { return edges[i][2] < edges[j][2] })

	type neighbor struct{ to, edgeID int }
	graph := make([][]neighbor, n)
	dfn := make([]int, n)
	timestamp := 0
	var tarjan func(int, int) int
	tarjan = func(v, pid int) int {
		timestamp++
		dfn[v] = timestamp
		lowV := timestamp
		for _, e := range graph[v] {
			if w := e.to; dfn[w] == 0 {
				lowW := tarjan(w, e.edgeID)
				if lowW > dfn[v] {
					edgeType[e.edgeID] = 1
				}
				lowV = findCriticalAndPseudoCriticalEdgesUnionMin(lowV, lowW)
			} else if e.edgeID != pid {
				lowV = findCriticalAndPseudoCriticalEdgesUnionMin(lowV, dfn[w])
			}
		}
		return lowV
	}

	uf := newFindCriticalAndPseudoCriticalEdgesUnion(n)
	for i := 0; i < m; {
		var vs []int
		for weight := edges[i][2]; i < m && edges[i][2] == weight; i++ {
			e := edges[i]
			v, w, edgeID := uf.find(e[0]), uf.find(e[1]), e[3]
			if v != w {
				graph[v] = append(graph[v], neighbor{w, edgeID})
				graph[w] = append(graph[w], neighbor{v, edgeID})
				vs = append(vs, v, w)
			} else {
				edgeType[edgeID] = -1
			}
		}
		for _, v := range vs {
			if dfn[v] == 0 {
				tarjan(v, -1)
			}
		}
		for j := 0; j < len(vs); j += 2 {
			v, w := vs[j], vs[j+1]
			uf.union(v, w)
			graph[v] = nil
			graph[w] = nil
			dfn[v] = 0
			dfn[w] = 0
		}
	}

	var keyEdges, pseudoKeyEdges []int
	for i, tp := range edgeType {
		if tp == 0 {
			pseudoKeyEdges = append(pseudoKeyEdges, i)
		} else if tp == 1 {
			keyEdges = append(keyEdges, i)
		}
	}
	return [][]int{keyEdges, pseudoKeyEdges}
}

func findCriticalAndPseudoCriticalEdgesUnionMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
