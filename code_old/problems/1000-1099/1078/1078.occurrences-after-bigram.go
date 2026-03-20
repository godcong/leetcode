package _1078

import (
	"strings"
)

func findOcurrences(text string, first string, second string) []string {
	words := strings.Split(text, " ")
	ans := make([]string, 0)
	for i := 2; i < len(words); i++ {
		if words[i-2] == first && words[i-1] == second {
			ans = append(ans, words[i])
		}
	}
	return ans
}
