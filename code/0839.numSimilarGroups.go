package code

/*
839. 相似字符串组
如果交换字符串 X 中的两个不同位置的字母，使得它和字符串 Y 相等，那么称 X 和 Y 两个字符串相似。如果这两个字符串本身是相等的，那它们也是相似的。

例如，"tars" 和 "rats" 是相似的 (交换 0 与 2 的位置)； "rats" 和 "arts" 也是相似的，但是 "star" 不与 "tars"，"rats"，或 "arts" 相似。

总之，它们通过相似性形成了两个关联组：{"tars", "rats", "arts"} 和 {"star"}。注意，"tars" 和 "arts" 是在同一组中，即使它们并不相似。形式上，对每个组而言，要确定一个单词在组中，只需要这个词和该组中至少一个单词相似。

给你一个字符串列表 strs。列表中的每个字符串都是 strs 中其它所有字符串的一个字母异位词。请问 strs 中有多少个相似字符串组？



示例 1：

输入：strs = ["tars","rats","arts","star"]
输出：2
示例 2：

输入：strs = ["omv","ovm"]
输出：1


提示：

1 <= strs.length <= 100
1 <= strs[i].length <= 1000
sum(strs[i].length) <= 2 * 104
strs[i] 只包含小写字母。
strs 中的所有单词都具有相同的长度，且是彼此的字母异位词。


备注：

      字母异位词（anagram），一种把某个字符串的字母的位置（顺序）加以改换所形成的新词。
*/
func numSimilarGroups(strs []string) int {
	m := map[string]bool{}
	for _, s := range strs {
		m[s] = true
	}
	i, ss := 0, make([]string, len(m))
	for s, _ := range m {
		ss[i], i = s, i+1
	}

	l := len(ss)
	if l <= 1 {
		return 1
	}
	au := (&numSimilarGroupsUnion{}).Init(l)
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			r1, r2 := au.GetRoot(i), au.GetRoot(j)
			if r1 == r2 {
				continue
			}
			sa, sb, c := ss[i], ss[j], 0
			for k := 0; k < len(sa); k++ {
				if sa[k] != sb[k] {
					c++
					if c > 2 {
						break
					}
				}
			}
			if c <= 2 {
				au.Set(r1, r2)
			}
		}
	}
	o := 0
	for i := 0; i < l; i++ {
		if au.Get(i) == -1 {
			o++
		}
	}
	return o
}

type numSimilarGroupsUnion struct {
	arr []int
}

func (au *numSimilarGroupsUnion) Init(l int) *numSimilarGroupsUnion {
	au.arr = make([]int, l)
	for i := 0; i < l; i++ {
		au.arr[i] = -1
	}
	return au
}

func (au *numSimilarGroupsUnion) Get(i int) int {
	return au.arr[i]
}

func (au *numSimilarGroupsUnion) Set(i, v int) {
	au.arr[i] = v
}

func (au *numSimilarGroupsUnion) GetRoot(i int) int {
	for {
		r := au.arr[i]
		if r == i || r == -1 {
			return i
		}
		i = r
	}
}
