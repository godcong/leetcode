package _0524

import (
	"sort"
)

func findLongestWord(s string, dictionary []string) string {
	sort.Slice(dictionary, func(i, j int) bool {
		a, b := dictionary[i], dictionary[j]
		return len(a) > len(b) || len(a) == len(b) && a < b
	})

	for _, t := range dictionary {
		i := 0
		for j := range s {
			if s[j] == t[i] {
				i++
				if i == len(t) {
					return t
				}
			}
		}

	}

	return ""
}
