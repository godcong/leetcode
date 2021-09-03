package _InterviewQuestion_17_14

func smallestK(arr []int, k int) []int {
	if len(arr) == 0 || k == 0 {
		return nil
	}
	var partition = func(i, j int) int {
		tmp := arr[i]
		for i < j {
			for ; i < j && arr[j] >= tmp; j-- {
			}
			arr[i] = arr[j]
			for ; i < j && arr[i] <= tmp; i++ {
			}
			arr[j] = arr[i]
		}
		arr[i] = tmp
		return i
	}
	i, j := 0, len(arr)-1
	k--
	idx := partition(i, j)
	for idx != k {
		if idx > k {
			j = idx - 1
			idx = partition(i, j)
		} else {
			i = idx + 1
			idx = partition(i, j)
		}
	}
	return arr[:k+1]
}
