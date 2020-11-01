package leetcode

import "strings"

/*
140. 单词拆分 II
给定一个非空字符串 s 和一个包含非空单词列表的字典 wordDict，在字符串中增加空格来构建一个句子，使得句子中所有的单词都在词典中。返回所有这些可能的句子。

说明：

分隔时可以重复使用字典中的单词。
你可以假设字典中没有重复的单词。
示例 1：

输入:
s = "catsanddog"
wordDict = ["cat", "cats", "and", "sand", "dog"]
输出:
[
  "cats and dog",
  "cat sand dog"
]
示例 2：

输入:
s = "pineapplepenapple"
wordDict = ["apple", "pen", "applepen", "pine", "pineapple"]
输出:
[
  "pine apple pen apple",
  "pineapple pen apple",
  "pine applepen apple"
]
解释: 注意你可以重复使用字典中的单词。
示例 3：

输入:
s = "catsandog"
wordDict = ["cats", "dog", "sand", "and", "cat"]
输出:
[]
*/
func wordBreak(s string, wordDict []string) []string {
	var sentences []string
	wordSet := map[string]struct{}{}
	for _, w := range wordDict {
		wordSet[w] = struct{}{}
	}

	n := len(s)
	dp := make([][][]string, n)
	var backtrack func(index int) [][]string
	backtrack = func(index int) [][]string {
		if dp[index] != nil {
			return dp[index]
		}
		var wordsList [][]string
		for i := index + 1; i < n; i++ {
			word := s[index:i]
			if _, has := wordSet[word]; has {
				for _, nextWords := range backtrack(i) {
					wordsList = append(wordsList, append([]string{word}, nextWords...))
				}
			}
		}
		word := s[index:]
		if _, has := wordSet[word]; has {
			wordsList = append(wordsList, []string{word})
		}
		dp[index] = wordsList
		return wordsList
	}
	for _, words := range backtrack(0) {
		sentences = append(sentences, strings.Join(words, " "))
	}
	return sentences
}
