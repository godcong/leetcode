package _0345

import (
	"strings"
)

const reverse = "aeiouAEIOU"

func reverseVowels(s string) string {
	t := []byte(s)
	n := len(t)
	i, j := 0, n-1
	for i < j {
		for i < n && !strings.Contains(reverse, string(t[i])) {
			i++
		}
		for j > 0 && !strings.Contains(reverse, string(t[j])) {
			j--
		}
		if i < j {
			t[i], t[j] = t[j], t[i]
			i++
			j--
		}
	}
	return string(t)
}
