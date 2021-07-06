package _1418

import (
	"sort"
	"strconv"
)

func displayTable(orders [][]string) [][]string {
	// 从订单中获取餐品名称和桌号，统计每桌点餐数量
	nameSet := map[string]struct{}{}
	foodsCnt := map[int]map[string]int{}
	for _, order := range orders {
		id, _ := strconv.Atoi(order[1])
		food := order[2]
		nameSet[food] = struct{}{}
		if foodsCnt[id] == nil {
			foodsCnt[id] = map[string]int{}
		}
		foodsCnt[id][food]++
	}

	// 提取餐品名称，并按字母顺序排列
	n := len(nameSet)
	names := make([]string, 0, n)
	for name := range nameSet {
		names = append(names, name)
	}
	sort.Strings(names)

	// 提取桌号，并按餐桌桌号升序排列
	m := len(foodsCnt)
	ids := make([]int, 0, m)
	for id := range foodsCnt {
		ids = append(ids, id)
	}
	sort.Ints(ids)

	// 填写点菜展示表
	table := make([][]string, m+1)
	table[0] = make([]string, 1, n+1)
	table[0][0] = "Table"
	table[0] = append(table[0], names...)
	for i, id := range ids {
		cnt := foodsCnt[id]
		table[i+1] = make([]string, n+1)
		table[i+1][0] = strconv.Itoa(id)
		for j, name := range names {
			table[i+1][j+1] = strconv.Itoa(cnt[name])
		}
	}
	return table
}
