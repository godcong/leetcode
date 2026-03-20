package _0187

var _map = []int{'A': 0, 'C': 1, 'G': 2, 'T': 3}

func findRepeatedDnaSequences(s string) []string {
	ans := make([]string, 0)
	n := len(s)
	if n <= 10 {
		return ans
	}
	h := 0
	for _, b := range s[:10] {
		h = h<<2 | _map[b]
	}
	c := [1 << 20]int8{}
	c[h] = 1
	for i := 10; i < n; i++ {
		h = (h<<2 | _map[s[i]]) % (1 << 20)
		if c[h] < 2 {
			if c[h] == 1 {
				ans = append(ans, s[i-9:i+1])
			}
			c[h]++
		}
	}
	return ans
}
