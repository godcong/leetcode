package _0825

const mx = 121

func numFriendRequests(ages []int) int {
	var cnt, pre [mx]int
	for _, age := range ages {
		cnt[age]++
	}
	for i := 1; i < mx; i++ {
		pre[i] = pre[i-1] + cnt[i]
	}

	var res int
	for i := 15; i < mx; i++ {
		if cnt[i] > 0 {
			res += cnt[i] * (pre[i] - pre[i/2+7] - 1)
		}
	}
	return res
}
