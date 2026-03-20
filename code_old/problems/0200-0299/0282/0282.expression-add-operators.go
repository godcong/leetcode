package _0282

func addOperators(num string, target int) []string {
	n := len(num)
	ans := make([]string, 0)
	var backtrack func(expr []byte, i, res, mul int)
	backtrack = func(expr []byte, i, res, mul int) {
		if i == n {
			if res == target {
				ans = append(ans, string(expr))
			}
			return
		}
		signIndex := len(expr)
		if i > 0 {
			expr = append(expr, 0)
		}
		for j, val := i, 0; j < n && (j == i || num[i] != '0'); j++ {
			val = val*10 + int(num[j]-'0')
			expr = append(expr, num[j])
			if i == 0 {
				backtrack(expr, j+1, val, val)
			} else {
				expr[signIndex] = '+'
				backtrack(expr, j+1, res+val, val)
				expr[signIndex] = '-'
				backtrack(expr, j+1, res-val, -val)
				expr[signIndex] = '*'
				backtrack(expr, j+1, res-mul+mul*val, mul*val)
			}
		}
	}
	backtrack(make([]byte, 0, n*2-1), 0, 0, 0)
	return ans
}
