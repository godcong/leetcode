package code

/*
57. 插入区间
给出一个无重叠的 ，按照区间起始端点排序的区间列表。

在列表中插入一个新的区间，你需要确保列表中的区间仍然有序且不重叠（如果有必要的话，可以合并区间）。



示例 1：

输入：intervals = [[1,3],[6,9]], newInterval = [2,5]
输出：[[1,5],[6,9]]
示例 2：

输入：intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
输出：[[1,2],[3,10],[12,16]]
解释：这是因为新的区间 [4,8] 与 [3,5],[6,7],[8,10] 重叠。


注意：输入类型已在 2019 年 4 月 15 日更改。请重置为默认代码定义以获取新的方法签名。
*/
func insert(intervals [][]int, newInterval []int) [][]int {
	var ret [][]int
	left, right := newInterval[0], newInterval[1]
	merged := false
	for _, interval := range intervals {
		if interval[0] > right {
			if !merged {
				ret = append(ret, []int{left, right})
				merged = true
			}
			ret = append(ret, interval)
		} else if interval[1] < left {
			ret = append(ret, interval)
		} else {
			left = inertMin(left, interval[0])
			right = insertMax(right, interval[1])
		}
	}
	if !merged {
		ret = append(ret, []int{left, right})
	}
	return ret
}

func insertMax(l, r int) int {
	if l > r {
		return l
	}
	return r
}

func inertMin(l, r int) int {
	if l < r {
		return l
	}
	return r
}
