package _0705

/*
705. 设计哈希集合
不使用任何内建的哈希表库设计一个哈希集合（HashSet）。

实现 MyHashSet 类：

void add(key) 向哈希集合中插入值 key 。
bool contains(key) 返回哈希集合中是否存在这个值 key 。
void remove(key) 将给定值 key 从哈希集合中删除。如果哈希集合中没有这个值，什么也不做。

示例：

输入：
["MyHashSet", "add", "add", "contains", "contains", "add", "contains", "remove", "contains"]
[[], [1], [2], [1], [3], [2], [2], [2], [2]]
输出：
[null, null, null, true, false, null, true, null, false]

解释：
MyHashSet myHashSet = new MyHashSet();
myHashSet.add(1);      // set = [1]
myHashSet.add(2);      // set = [1, 2]
myHashSet.contains(1); // 返回 True
myHashSet.contains(3); // 返回 False ，（未找到）
myHashSet.add(2);      // set = [1, 2]
myHashSet.contains(2); // 返回 True
myHashSet.remove(2);   // set = [1]
myHashSet.contains(2); // 返回 False ，（已移除）


提示：

0 <= key <= 106
最多调用 104 次 add、remove 和 contains 。


进阶：你可以不使用内建的哈希集合库解决此问题吗？
*/
type MyHashSet struct {
	bitset []uint64
}


/** Initialize your data structure here. */
func Constructor() MyHashSet {
	return MyHashSet{bitset: []uint64{}}
}


func (s *MyHashSet) Add(key int)  {
	bit := key % 64
	length := key /64
	for i := len(s.bitset); i <= length; i++ {
		s.bitset = append(s.bitset, 0)
	}
	s.bitset[length] = s.bitset[length] | (1 << bit)
}


func (s *MyHashSet) Remove(key int)  {
	bit := key % 64
	length := key /64
	if length >= len(s.bitset) {
		return
	}
	s.bitset[length] = s.bitset[length] & ^(1 << bit)
}


/** Returns true if this set contains the specified element */
func (s *MyHashSet) Contains(key int) bool {
	bit := key % 64
	length := key /64
	if length >= len(s.bitset) {
		return false
	}
	return s.bitset[length] & (1 << bit) != 0
}
