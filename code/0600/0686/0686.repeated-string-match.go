package _0686

func repeatedStringMatch(a string, b string) int {
	an, bn := len(a), len(b)
	index := strStr(a, b)
	if index == -1 {
		return -1
	}
	if an-index >= bn {
		return 1
	}
	return (bn+index-an-1)/an + 2
}

func strStr(haystack, needle string) int {
	n, m := len(haystack), len(needle)
	if m == 0 {
		return 0
	}
	pi := make([]int, m)
	for i, j := 1, 0; i < m; i++ {
		for j > 0 && needle[i] != needle[j] {
			j = pi[j-1]
		}
		if needle[i] == needle[j] {
			j++
		}
		pi[i] = j
	}
	for i, j := 0, 0; i-j < n; i++ {
		for j > 0 && haystack[i%n] != needle[j] {
			j = pi[j-1]
		}
		if haystack[i%n] == needle[j] {
			j++
		}
		if j == m {
			return i - m + 1
		}
	}
	return -1
}
