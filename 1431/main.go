package main

import "fmt"

func kidsWithCandies(candies []int, extraCandies int) []bool {
	max := candies[0]
	size := len(candies)
	for i := 1; i < len(candies); i++ {
		if candies[i] > max {
			max = candies[i]
		}
	}
	max -= extraCandies
	ret := make([]bool, size)
	for i := 0; i < size; i++ {
		if candies[i] >= max {
			ret[i] = true
		}
	}
	return ret
}

func main() {
	fmt.Println("candies = [2,3,5,1,3], extraCandies = 3", kidsWithCandies([]int{2, 3, 5, 1, 3}, 3))
	fmt.Println("candies = [4,2,1,1,2], extraCandies = 1", kidsWithCandies([]int{4, 2, 1, 1, 2}, 1))
	fmt.Println("candies = [12,1,12], extraCandies = 10", kidsWithCandies([]int{12, 1, 12}, 10))
}
