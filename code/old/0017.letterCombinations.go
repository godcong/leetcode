package old

/*
17. 电话号码的字母组合
给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。

给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。



示例:

输入："23"
输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
说明:
尽管上面的答案是按字典序排列的，但是你可以任意选择答案输出的顺序。
*/
var letterCombinationsCharList = [8][]byte{
	{'a', 'b', 'c'},
	{'d', 'e', 'f'},
	{'g', 'h', 'i'},
	{'j', 'k', 'l'},
	{'m', 'n', 'o'},
	{'p', 'q', 'r', 's'},
	{'t', 'u', 'v'},
	{'w', 'x', 'y', 'z'},
}

func letterCombinations(digits string) []string {
	size := byte(len(digits))
	if size == 0 {
		return nil
	}
	var ret []string
	var letterCombinationsRecursion func(pre []byte, idx byte) []string
	letterCombinationsRecursion = func(pre []byte, idx byte) []string {
		for i := range letterCombinationsCharList[digits[idx]-'2'] {
			pre[idx] = letterCombinationsCharList[digits[idx]-'2'][i]
			if idx+1 < size {
				letterCombinationsRecursion(pre, idx+1)
			} else {
				ret = append(ret, string(pre))
			}
		}
		return ret
	}
	b := make([]byte, size)
	return letterCombinationsRecursion(b, 0)
}
