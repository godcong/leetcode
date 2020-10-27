package leetcode

import (
	"strconv"
	"strings"
)

func strToIntArray(nums string) []int {
	sta := 0
	end := len(nums)
	if strings.Index(nums, "[") == 0 {
		sta += 1
	}

	if strings.Index(nums, "]") == end-1 {
		end -= 1
	}

	var sNums []string
	nums = nums[sta:end]
	if nums == "" {
		return nil
	} else {
		sNums = strings.Split(nums, ",")
	}

	size := len(sNums)
	nodes := make([]int, size)
	for i, num := range sNums {
		atoi, err := strconv.Atoi(num)
		if err != nil {
			panic(err)
		}
		nodes[i] = atoi
	}
	return nodes
}
