package _1109

func corpFlightBookings(bookings [][]int, n int) []int {
	nums := make([]int, n)
	for _, booking := range bookings {
		nums[booking[0]-1] += booking[2]
		if booking[1] < n {
			nums[booking[1]] -= booking[2]
		}
	}
	for i := 1; i < n; i++ {
		nums[i] += nums[i-1]
	}
	return nums
}
