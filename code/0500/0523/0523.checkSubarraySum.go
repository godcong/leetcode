package _0523

func checkSubarraySum(nums []int, k int) bool {
	sumSet := map[int]int{}
	sumSet[0] = -1
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if k != 0 {
			sum = sum % k
		}
		index, extist := sumSet[sum]
		if extist && i-index > 1 {
			return true
		} else if !extist {
			sumSet[sum] = i
		}
	}
	return false
}
