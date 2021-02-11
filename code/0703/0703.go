package _703

import "math"

/*
703. 数据流中的第 K 大元素
设计一个找到数据流中第 k 大元素的类（class）。注意是排序后的第 k 大元素，不是第 k 个不同的元素。

请实现 KthLargest 类：

KthLargest(int k, int[] nums) 使用整数 k 和整数流 nums 初始化对象。
int add(int val) 将 val 插入数据流 nums 后，返回当前数据流中第 k 大的元素。


示例：

输入：
["KthLargest", "add", "add", "add", "add", "add"]
[[3, [4, 5, 8, 2]], [3], [5], [10], [9], [4]]
输出：
[null, 4, 5, 5, 8, 8]

解释：
KthLargest kthLargest = new KthLargest(3, [4, 5, 8, 2]);
kthLargest.add(3);   // return 4
kthLargest.add(5);   // return 5
kthLargest.add(10);  // return 5
kthLargest.add(9);   // return 8
kthLargest.add(4);   // return 8


提示：
1 <= k <= 104
0 <= nums.length <= 104
-104 <= nums[i] <= 104
-104 <= val <= 104
最多调用 add 方法 104 次
题目数据保证，在查找第 k 大元素时，数组中至少有 k 个元素
*/
type KthLargest struct {
	k    int
	nums []int // 切片中只放置k个元素
}

func Constructor(k int, nums []int) KthLargest {
	kthLargest := KthLargest{}
	kthLargest.k = k
	kthLargest.nums = []int{math.MinInt32}
	for _, num := range nums {
		kthLargest.Add(num)
	}
	return kthLargest
}

func (this *KthLargest) Add(val int) int {
	if len(this.nums) == 0 {
		this.actualAdd(0, val)
		return this.nums[0]
	}
	if len(this.nums) < this.k {
		// 长度不够时, 直接add
		this.actualAdd(len(this.nums), val)
	} else {
		// 数据量多于k个, 如果val > 堆顶元素(第k大的元素)
		// 则取出栈顶元素, 再将val放进去
		// 否则, 直接丢弃val元素
		if val > this.nums[0] {
			// 取最后一个元素往下冒泡
			this.siftDown(0, this.nums[len(this.nums) - 1])
			// 砍掉最后一个元素
			this.nums = this.nums[: len(this.nums) - 1]
			// 添加新元素
			this.actualAdd(len(this.nums), val)
		}
	}
	return this.nums[0]
}

func (this *KthLargest) actualAdd(k, key int) {
	this.nums = append(this.nums, key)
	// find parent
	for k > 0 {
		pntIdx := (k - 1) >> 1
		e := this.nums[pntIdx]
		if key >= e {
			break
		}
		this.nums[k] = e
		k = pntIdx
	}
	this.nums[k] = key
}

func (this *KthLargest) siftUp(pntIdx, idx int) {
	pntNum := this.nums[pntIdx]
	num := this.nums[idx]
	if pntNum > num {
		this.nums[pntIdx] = num
		this.nums[idx] = pntNum
		idx = pntIdx
		pntIdx = (idx - 1) / 2
		if pntIdx >= 0 && this.nums[pntIdx] > this.nums[idx] {
			this.siftUp(pntIdx, idx)
		}
	}
}

func (this *KthLargest) siftDown(k, key int) {
	size := len(this.nums)
	half := size >> 1
	for k < half {
		// 左子节点索引
		childIdx := (k << 1) + 1
		// 左子节点
		c := this.nums[childIdx]
		// 右子节点索引
		rightIdx := childIdx + 1
		if rightIdx < size && c > this.nums[rightIdx] {
			// 取出左右子节点中较小的节点索引
			childIdx = rightIdx
			// 较小的节点
			c = this.nums[childIdx]
		}
		// 如果key小于 最小子节点, 直接跳出
		if key <= c {
			break
		}
		// 将key, 左子节点, 右子节点中最小的值, 至于目标k处
		this.nums[k] = c
		k = childIdx
	}
	this.nums[k] = key
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(length, nums);
 * param_1 := obj.Add(val);
 */
