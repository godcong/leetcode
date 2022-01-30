package _0884

import (
	"strings"
)

func uncommonFromSentences(s1 string, s2 string) []string {
	freq := make(map[string]int)

	insert := func(s string) {
		words := strings.Split(s, " ")
		for _, word := range words {
			freq[word]++
		}
	}

	insert(s1)
	insert(s2)

	ans := make([]string, 0)
	for word, occ := range freq {
		if occ == 1 {
			ans = append(ans, word)
		}
	}
	return ans
}
