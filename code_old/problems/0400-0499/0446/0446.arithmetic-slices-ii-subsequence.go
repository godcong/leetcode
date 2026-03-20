package _0446

func numberOfArithmeticSlices(nums []int) int {
	f := make([]map[int]int, len(nums))
	ans := 0
	for i, x := range nums {
		f[i] = map[int]int{}
		for j, y := range nums[:i] {
			d := x - y
			cnt := f[j][d]
			ans += cnt
			f[i][d] += cnt + 1
		}
	}
	return ans
}
