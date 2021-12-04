package _0383

func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}
	cnt := [26]int{}
	for _, ch := range magazine {
		cnt[ch-'a']++
	}
	for _, ch := range ransomNote {
		cnt[ch-'a']--
		if cnt[ch-'a'] < 0 {
			return false
		}
	}
	return true
}
