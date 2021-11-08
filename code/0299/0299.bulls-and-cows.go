package _0299

import (
	"fmt"
)

func getHint(secret string, guess string) string {
	var bulls, cows int
	s, g := make([]int, 10), make([]int, 10)
	for i := 0; i < len(secret); i++ {
		if secret[i] == guess[i] {
			bulls++
		} else {
			s[secret[i]-'0']++
			g[guess[i]-'0']++
		}
	}
	for i := 0; i < 10; i++ {
		cows += min(s[i], g[i])
	}
	return fmt.Sprintf("%dA%dB", bulls, cows)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
