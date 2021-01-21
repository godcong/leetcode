package code

import "sort"

/*
721. 账户合并
给定一个列表 accounts，每个元素 accounts[i] 是一个字符串列表，其中第一个元素 accounts[i][0] 是 名称 (name)，其余元素是 emails 表示该账户的邮箱地址。

现在，我们想合并这些账户。如果两个账户都有一些共同的邮箱地址，则两个账户必定属于同一个人。请注意，即使两个账户具有相同的名称，它们也可能属于不同的人，因为人们可能具有相同的名称。一个人最初可以拥有任意数量的账户，但其所有账户都具有相同的名称。

合并账户后，按以下格式返回账户：每个账户的第一个元素是名称，其余元素是按字符 ASCII 顺序排列的邮箱地址。账户本身可以以任意顺序返回。



示例 1：

输入：
accounts = [["John", "johnsmith@mail.com", "john00@mail.com"], ["John", "johnnybravo@mail.com"], ["John", "johnsmith@mail.com", "john_newyork@mail.com"], ["Mary", "mary@mail.com"]]
输出：
[["John", 'john00@mail.com', 'john_newyork@mail.com', 'johnsmith@mail.com'],  ["John", "johnnybravo@mail.com"], ["Mary", "mary@mail.com"]]
解释：
第一个和第三个 John 是同一个人，因为他们有共同的邮箱地址 "johnsmith@mail.com"。
第二个 John 和 Mary 是不同的人，因为他们的邮箱地址没有被其他帐户使用。
可以以任何顺序返回这些列表，例如答案 [['Mary'，'mary@mail.com']，['John'，'johnnybravo@mail.com']，
['John'，'john00@mail.com'，'john_newyork@mail.com'，'johnsmith@mail.com']] 也是正确的。


提示：

accounts的长度将在[1，1000]的范围内。
accounts[i]的长度将在[1，10]的范围内。
accounts[i][j]的长度将在[1，30]的范围内。
*/
func accountsMerge(accounts [][]string) [][]string {
	email2AccountID := make(map[string]int)
	parent := make([]int, 1001)
	ans := make(map[int][]string)
	for i := 0; i < len(parent); i++ {
		parent[i] = -1
	}

	for i := 0; i < len(accounts); i++ {

		for j := 1; j < len(accounts[i]); j++ {

			if countID, ok := email2AccountID[accounts[i][j]]; ok {

				xPos := countID

				yPos := i

				accountsMergeUnion(parent, yPos, xPos)
			} else {

				email2AccountID[accounts[i][j]] = i
			}
		}
	}

	for email, countID := range email2AccountID {
		root := accountsMergeFind(parent, countID)
		ans[root] = append(ans[root], email)
	}

	var res [][]string
	for countID, emails := range ans {

		sort.Slice(emails, func(i, j int) bool {
			return emails[i] < emails[j]
		})
		res = append(res, append([]string{accounts[countID][0]}, emails...))
	}
	return res
}

func accountsMergeFind(parent []int, value int) int {
	for parent[value] != -1 {
		value = parent[value]
	}
	return value
}

func accountsMergeUnion(parent []int, x, y int) {
	xRoot := accountsMergeFind(parent, x)
	yRoot := accountsMergeFind(parent, y)
	if xRoot != yRoot {
		parent[xRoot] = yRoot
	}
}
