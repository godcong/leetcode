package _1588

func sumOddLengthSubarrays(arr []int) int {
	res, lens := 0, len(arr)
	for i := 1; i < lens; i++ {
		arr[i] += arr[i-1]
	}
	for i := 1; i <= lens; i += 2 {
		a, b := 0, i
		res += arr[b-1]
		for b < lens {
			res += arr[b] - arr[a]
			a++
			b++
		}
	}
	return res
}
