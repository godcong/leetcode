package code

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
	if s == "" || len(s) == 0 || wordDict == nil || len(wordDict) == 0 {
		return nil
	}
	var res []string
	set := [26]bool{}

	for _, v := range strings.Join(wordDict, "") {
		set[v-'a'] = true
	}

	for _, v := range s {
		if !set[v-'a'] {
			return res
		}
	}
	var wordBreakDFS func(nowString string, index int)
	max := len(s)
	wordBreakDFS = func(nowString string, index int) {

		if index >= max {
			res = append(res, nowString[1:])
			return
		}

		for _, word := range wordDict {
			lenWord := len(word)

			if index+lenWord <= len(s) && s[index:index+lenWord] == word {
				wordBreakDFS(nowString+" "+word, index+lenWord)
			}
		}
	}
	wordBreakDFS("", 0)
	return res
}
