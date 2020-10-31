package leetcode

import (
	"fmt"
	"math/rand"
)

/*
381. O(1) 时间插入、删除和获取随机元素 - 允许重复
设计一个支持在平均 时间复杂度 O(1) 下， 执行以下操作的数据结构。

注意: 允许出现重复元素。

insert(val)：向集合中插入元素 val。
remove(val)：当 val 存在时，从集合中移除一个 val。
getRandom：从现有集合中随机获取一个元素。每个元素被返回的概率应该与其在集合中的数量呈线性相关。
示例:

// 初始化一个空的集合。
RandomizedCollection collection = new RandomizedCollection();

// 向集合中插入 1 。返回 true 表示集合不包含 1 。
collection.insert(1);

// 向集合中插入另一个 1 。返回 false 表示集合包含 1 。集合现在包含 [1,1] 。
collection.insert(1);

// 向集合中插入 2 ，返回 true 。集合现在包含 [1,1,2] 。
collection.insert(2);

// getRandom 应当有 2/3 的概率返回 1 ，1/3 的概率返回 2 。
collection.getRandom();

// 从集合中删除 1 ，返回 true 。集合现在包含 [1,2] 。
collection.remove(1);

// getRandom 应有相同概率返回 1 和 2 。
collection.getRandom();
*/
type RandomizedCollection struct {
	vals map[int]map[int]bool
	nums []int
}

/** Initialize your data structure here. */
func Constructor() RandomizedCollection {
	return RandomizedCollection{
		vals: make(map[int]map[int]bool),
	}
}

/** Inserts a value to the collection. Returns true if the collection did not already contain the specified element. */
func (this *RandomizedCollection) Insert(val int) bool {
	idx, b := this.vals[val]
	if !b {
		idx = map[int]bool{}
		this.vals[val] = idx
	}
	idx[len(this.nums)] = true
	this.nums = append(this.nums, val)
	return b
}

/** Removes a value from the collection. Returns true if the collection contained the specified element. */
func (this *RandomizedCollection) Remove(val int) bool {
	idx, has := this.vals[val]
	if !has {
		return false
	}
	var i int
	for id := range idx {
		i = id
		break
	}
	last := len(this.nums) - 1
	this.nums[i] = this.nums[last]
	delete(idx, i)
	delete(this.vals[this.nums[i]], last)
	if i < last {
		this.vals[this.nums[i]][i] = true
	}
	if len(idx) == 0 {
		delete(this.vals, val)
	}
	this.nums = this.nums[:last]
	return true
}

/** Get a random element from the collection. */
func (this *RandomizedCollection) GetRandom() int {
	fmt.Println("GetRandom", this.nums)
	return this.nums[rand.Intn(len(this.nums))]
}

/**
 * Your RandomizedCollection object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
