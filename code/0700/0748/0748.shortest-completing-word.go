package _0748

import (
	"unicode"
)

func shortestCompletingWord(licensePlate string, words []string) string {
	cnt := [26]int{}
	for _, ch := range licensePlate {
		if unicode.IsLetter(ch) {
			cnt[unicode.ToLower(ch)-'a']++
		}
	}
	var ans string
next:
	for _, word := range words {
		c := [26]int{}
		for _, ch := range word {
			c[ch-'a']++
		}
		for i := 0; i < 26; i++ {
			if c[i] < cnt[i] {
				continue next
			}
		}
		if ans == "" || len(word) < len(ans) {
			ans = word
		}
	}
	return ans
}
