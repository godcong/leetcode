package old

/*
1579. 保证图可完全遍历
Alice 和 Bob 共有一个无向图，其中包含 n 个节点和 3  种类型的边：

类型 1：只能由 Alice 遍历。
类型 2：只能由 Bob 遍历。
类型 3：Alice 和 Bob 都可以遍历。
给你一个数组 edges ，其中 edges[i] = [typei, ui, vi] 表示节点 ui 和 vi 之间存在类型为 typei 的双向边。请你在保证图仍能够被 Alice和 Bob 完全遍历的前提下，找出可以删除的最大边数。如果从任何节点开始，Alice 和 Bob 都可以到达所有其他节点，则认为图是可以完全遍历的。

返回可以删除的最大边数，如果 Alice 和 Bob 无法完全遍历图，则返回 -1 。



示例 1：



输入：n = 4, edges = [[3,1,2],[3,2,3],[1,1,3],[1,2,4],[1,1,2],[2,3,4]]
输出：2
解释：如果删除 [1,1,2] 和 [1,1,3] 这两条边，Alice 和 Bob 仍然可以完全遍历这个图。再删除任何其他的边都无法保证图可以完全遍历。所以可以删除的最大边数是 2 。
示例 2：



输入：n = 4, edges = [[3,1,2],[3,2,3],[1,1,4],[2,1,4]]
输出：0
解释：注意，删除任何一条边都会使 Alice 和 Bob 无法完全遍历这个图。
示例 3：



输入：n = 4, edges = [[3,2,3],[1,1,2],[2,3,4]]
输出：-1
解释：在当前图中，Alice 无法从其他节点到达节点 4 。类似地，Bob 也不能达到节点 1 。因此，图无法完全遍历。


提示：

1 <= n <= 10^5
1 <= edges.length <= min(10^5, 3 * n * (n-1) / 2)
edges[i].length == 3
1 <= edges[i][0] <= 3
1 <= edges[i][1] < edges[i][2] <= n
所有元组 (typei, ui, vi) 互不相同
*/
func maxNumEdgesToRemove(n int, edges [][]int) int {
	au := (&maxNumEdgesToRemoveUnion{}).Init(n + 1)
	o, c1 := 0, n
	for _, e := range edges {
		if e[0] == 3 {
			if !au.Merge(e[1], e[2]) {
				o++
			} else {
				c1--
			}
		}
	}
	au1 := au.Copy()
	c2 := c1
	for _, e := range edges {
		if e[0] == 1 {
			if !au.Merge(e[1], e[2]) {
				o++
			} else {
				c1--
			}
		} else if e[0] == 2 {
			if !au1.Merge(e[1], e[2]) {
				o++
			} else {
				c2--
			}
		}
	}
	if c1 != 1 || c2 != 1 {
		return -1
	}
	return o
}

type maxNumEdgesToRemoveUnion struct {
	arr []int
}

func (au *maxNumEdgesToRemoveUnion) Init(l int) *maxNumEdgesToRemoveUnion {
	au.arr = make([]int, l)
	for i := 0; i < l; i++ {
		au.arr[i] = i
	}
	return au
}

func (au *maxNumEdgesToRemoveUnion) Copy() *maxNumEdgesToRemoveUnion {
	arr := make([]int, len(au.arr))
	copy(arr, au.arr)
	return &maxNumEdgesToRemoveUnion{arr}
}

func (au *maxNumEdgesToRemoveUnion) Set(i, v int) {
	au.arr[i] = v
}

func (au *maxNumEdgesToRemoveUnion) GetRoot(i int) int {
	r := au.arr[i]
	if r == i {
		return i
	}
	r = au.GetRoot(r)
	au.arr[i] = r
	return r
}

func (au *maxNumEdgesToRemoveUnion) Merge(i, j int) bool {
	r1, r2 := au.GetRoot(i), au.GetRoot(j)
	if r1 != r2 {
		au.Set(r1, r2)
		return true
	}
	return false
}
