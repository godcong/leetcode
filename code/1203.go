package code

/*
1203. 项目管理
公司共有 n 个项目和  m 个小组，每个项目要不无人接手，要不就由 m 个小组之一负责。

group[i] 表示第 i 个项目所属的小组，如果这个项目目前无人接手，那么 group[i] 就等于 -1。（项目和小组都是从零开始编号的）小组可能存在没有接手任何项目的情况。

请你帮忙按要求安排这些项目的进度，并返回排序后的项目列表：

同一小组的项目，排序后在列表中彼此相邻。
项目之间存在一定的依赖关系，我们用一个列表 beforeItems 来表示，其中 beforeItems[i] 表示在进行第 i 个项目前（位于第 i 个项目左侧）应该完成的所有项目。
如果存在多个解决方案，只需要返回其中任意一个即可。如果没有合适的解决方案，就请返回一个 空列表 。



示例 1：



输入：n = 8, m = 2, group = [-1,-1,1,0,0,1,0,-1], beforeItems = [[],[6],[5],[6],[3,6],[],[],[]]
输出：[6,3,4,1,5,2,0,7]
示例 2：

输入：n = 8, m = 2, group = [-1,-1,1,0,0,1,0,-1], beforeItems = [[],[6],[5],[6],[3],[],[4],[]]
输出：[]
解释：与示例 1 大致相同，但是在排序后的列表中，4 必须放在 6 的前面。


提示：

1 <= m <= n <= 3 * 104
group.length == beforeItems.length == n
-1 <= group[i] <= m - 1
0 <= beforeItems[i].length <= n - 1
0 <= beforeItems[i][j] <= n - 1
i != beforeItems[i][j]
beforeItems[i] 不含重复元素
*/
func sortItems(n int, m int, group []int, beforeItems [][]int) []int {
	var ans []int
	groupItems := make([][]int, m+n) // groupItems[i] 表示第 i 个组负责的项目列表
	for i := range group {
		if group[i] == -1 {
			group[i] = m + i // 给不属于任何组的项目分配一个组
		}
		groupItems[group[i]] = append(groupItems[group[i]], i)
	}

	groupGraph := make([][]int, m+n) // 组间依赖关系
	groupDegree := make([]int, m+n)
	itemGraph := make([][]int, n) // 组内依赖关系
	itemDegree := make([]int, n)
	for cur, items := range beforeItems {
		curGroupID := group[cur]
		for _, pre := range items {
			preGroupID := group[pre]
			if preGroupID != curGroupID { // 不同组项目，确定组间依赖关系
				groupGraph[preGroupID] = append(groupGraph[preGroupID], curGroupID)
				groupDegree[curGroupID]++
			} else { // 同组项目，确定组内依赖关系
				itemGraph[pre] = append(itemGraph[pre], cur)
				itemDegree[cur]++
			}
		}
	}

	// 组间拓扑序
	items := make([]int, m+n)
	for i := range items {
		items[i] = i
	}
	groupOrders := sortItemsTopSort(groupGraph, groupDegree, items)
	if len(groupOrders) < len(items) {
		return nil
	}

	// 按照组间的拓扑序，依次求得各个组的组内拓扑序，构成答案
	for _, groupID := range groupOrders {
		items := groupItems[groupID]
		orders := sortItemsTopSort(itemGraph, itemDegree, items)
		if len(orders) < len(items) {
			return nil
		}
		ans = append(ans, orders...)
	}
	return ans
}
func sortItemsTopSort(graph [][]int, deg, items []int) (orders []int) {
	q := []int{}
	for _, i := range items {
		if deg[i] == 0 {
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		from := q[0]
		q = q[1:]
		orders = append(orders, from)
		for _, to := range graph[from] {
			deg[to]--
			if deg[to] == 0 {
				q = append(q, to)
			}
		}
	}
	return
}
