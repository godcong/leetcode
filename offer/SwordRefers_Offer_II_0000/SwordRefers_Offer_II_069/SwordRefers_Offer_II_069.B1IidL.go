package _SwordRefers_Offer_II_069

func peakIndexInMountainArray(arr []int) int {
	left := 0
	right := len(arr) - 1

	for left < right {

		mid := (left + right) / 2

		if arr[mid] > arr[mid-1] && arr[mid] > arr[mid+1] {
			return mid
		}

		if arr[mid] > arr[mid-1] {
			left = mid
		}

		if arr[mid] > arr[mid+1] {
			right = mid
		}
	}

	return 0
}
