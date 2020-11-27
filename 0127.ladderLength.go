package leetcode

/*
127. 单词接龙
给定两个单词（beginWord 和 endWord）和一个字典，找到从 beginWord 到 endWord 的最短转换序列的长度。转换需遵循如下规则：

每次转换只能改变一个字母。
转换过程中的中间单词必须是字典中的单词。
说明:

如果不存在这样的转换序列，返回 0。
所有单词具有相同的长度。
所有单词只由小写字母组成。
字典中不存在重复的单词。
你可以假设 beginWord 和 endWord 是非空的，且二者不相同。
示例 1:

输入:
beginWord = "hit",
endWord = "cog",
wordList = ["hot","dot","dog","lot","log","cog"]

输出: 5

解释: 一个最短转换序列是 "hit" -> "hot" -> "dot" -> "dog" -> "cog",
     返回它的长度 5。
示例 2:

输入:
beginWord = "hit"
endWord = "cog"
wordList = ["hot","dot","dog","lot","log"]

输出: 0

解释: endWord "cog" 不在字典中，所以无法进行转换。
*/
func ladderLength(beginWord string, endWord string, wordList []string) int {
	if ladderLengthIndexOf(endWord, wordList) == -1 {
		return 0
	}
	step := 0
	startQueue := []string{beginWord}
	endQueue := []string{endWord}
	used := make([]bool, len(wordList))

	for len(startQueue) > 0 {
		step++
		l := len(startQueue)
		for i := 0; i < l; i++ {
			for k := 0; k < len(endQueue); k++ {
				if ladderLengthHasOneDiff(startQueue[i], endQueue[k]) {
					return step + 1
				}
			}
			for j, w := range wordList {
				if !used[j] && ladderLengthHasOneDiff(startQueue[i], w) {
					startQueue = append(startQueue, w)
					used[j] = true
				}
			}
		}
		startQueue = startQueue[l:]
		if len(startQueue) > len(endQueue) {
			startQueue, endQueue = endQueue, startQueue
		}
	}
	return 0
}

func ladderLengthIndexOf(str string, bank []string) int {
	for i, s := range bank {
		if str == s {
			return i
		}
	}
	return -1
}

func ladderLengthHasOneDiff(x, y string) bool {
	count := 0
	for i := 0; i < len(x); i++ {
		if x[i] != y[i] {
			count++
		}
		if count > 1 {
			return false
		}
	}
	if count == 1 {
		return true
	}
	return false
}
