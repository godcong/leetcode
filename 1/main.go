package main

import "fmt"

func twoSum(nums []int, target int) []int {
	tmp := make(map[int]int)
	for idx, num := range nums {
		if v, exist := tmp[target-num]; exist {
			return []int{v, idx}
		}
		tmp[num] = idx
	}
	panic("unreachable")
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 18))
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 13))
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 26))

}
