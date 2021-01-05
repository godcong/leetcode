package code

import (
	"testing"
)

func TestRandomizedCollection(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			name: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 初始化一个空的集合。
			collection := Constructor()

			// 向集合中插入 1 。返回 true 表示集合不包含 1 。
			collection.Insert(1)

			// 向集合中插入另一个 1 。返回 false 表示集合包含 1 。集合现在包含 [1,1] 。
			collection.Insert(1)

			// 向集合中插入 2 ，返回 true 。集合现在包含 [1,1,2] 。
			collection.Insert(2)

			// getRandom 应当有 2/3 的概率返回 1 ，1/3 的概率返回 2 。
			collection.GetRandom()

			// 从集合中删除 1 ，返回 true 。集合现在包含 [1,2] 。
			collection.Remove(1)

			// getRandom 应有相同概率返回 1 和 2 。
			collection.GetRandom()

		})
	}
}
