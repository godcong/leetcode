package _0500

import (
	"unicode"
)

const rowIdx = "12210111011122000010020202"

func findWords(words []string) []string {
	var ans []string
next:
	for _, word := range words {
		idx := rowIdx[unicode.ToLower(rune(word[0]))-'a']
		for _, ch := range word[1:] {
			if rowIdx[unicode.ToLower(ch)-'a'] != idx {
				continue next
			}
		}
		ans = append(ans, word)
	}
	return ans
}
