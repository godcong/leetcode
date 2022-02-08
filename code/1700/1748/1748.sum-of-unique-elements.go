package _1748

func sumOfUnique(nums []int) int {
	cnt := map[int]int{}
	for _, num := range nums {
		cnt[num]++
	}
	ans := 0
	for num, c := range cnt {
		if c == 1 {
			ans += num
		}
	}
	return ans
}
