package _0852

func peakIndexInMountainArray(arr []int) int {
	for i := 1; ; i++ {
		if arr[i] > arr[i+1] {
			return i
		}
	}
}
