package _1414

func findMinFibonacciNumbers(k int) int {
	f := []int{1, 1}
	for f[len(f)-1] < k {
		f = append(f, f[len(f)-1]+f[len(f)-2])
	}
	ans := 0
	for i := len(f) - 1; k > 0; i-- {
		if k >= f[i] {
			k -= f[i]
			ans++
		}
	}
	return ans
}
