package code

import "sort"

/*
406. 根据身高重建队列
假设有打乱顺序的一群人站成一个队列。 每个人由一个整数对(h, k)表示，其中h是这个人的身高，k是排在这个人前面且身高大于或等于h的人数。 编写一个算法来重建这个队列。

注意：
总人数少于1100人。

示例

输入:
[[7,0], [4,4], [7,1], [5,0], [6,1], [5,2]]

输出:
[[5,0], [7,0], [5,2], [6,1], [4,4], [7,1]]
*/
func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		p, q := people[i], people[j]
		return p[0] > q[0] || p[0] == q[0] && p[1] < q[1]
	})
	var tmp []int
	for i := range people {
		tmp = people[i]
		for j := i; j > tmp[1]; j-- {
			people[j] = people[j-1]
		}
		people[tmp[1]] = tmp
	}
	return people
}
