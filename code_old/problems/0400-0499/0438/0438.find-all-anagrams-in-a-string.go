package _0438

func findAnagrams(s string, p string) []int {
	var ans []int
	sLen, pLen := len(s), len(p)
	if sLen < pLen {
		return ans
	}

	var sCount, pCount [26]int
	for i, ch := range p {
		sCount[s[i]-'a']++
		pCount[ch-'a']++
	}
	if sCount == pCount {
		ans = append(ans, 0)
	}

	for i, ch := range s[:sLen-pLen] {
		sCount[ch-'a']--
		sCount[s[i+pLen]-'a']++
		if sCount == pCount {
			ans = append(ans, i+1)
		}
	}
	return ans

}
