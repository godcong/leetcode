package _0397

func integerReplacement(n int) int {
	if n == 1 {
		return 0
	}
	if n%2 == 0 {
		return 1 + integerReplacement(n/2)
	}
	return 2 + min(integerReplacement(n/2), integerReplacement(n/2+1))
}

func min(replacement int, replacement2 int) int {
	if replacement < replacement2 {
		return replacement
	}
	return replacement2
}
