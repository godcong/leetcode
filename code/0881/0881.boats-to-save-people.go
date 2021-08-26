package _0881

func numRescueBoats(people []int, limit int) int {
	s, res := make([]int, limit+1), 0
	for _, v := range people {
		s[v]++
	}
	for i, j := 1, limit; i <= j; {
		if j+i > limit {
			res += s[j]
			j--
		} else if i == j {
			res += s[i]/2 + s[i]%2
			break
		} else {
			if s[j] > s[i] {
				res += s[i]
				s[j] -= s[i]
				i++
			} else if s[j] < s[i] {
				res += s[j]
				s[i] -= s[j]
				j--
			} else {
				res += s[i]
				i++
				j--
			}
		}
	}

	return res
}
