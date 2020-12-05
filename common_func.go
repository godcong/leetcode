package leetcode

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func fixBrackets(nums string) string {
	if nums == "" {
		return ""
	}

	sta := 0
	end := len(nums)
	if strings.Index(nums, "[") == 0 {
		sta += 1
	}

	if strings.LastIndex(nums, "]") == end-1 {
		end -= 1
	}

	return nums[sta:end]
}

func strToIntArray(nums string) []int {
	nums = fixBrackets(nums)
	var sNums []string
	if nums == "" {
		return nil
	} else {
		sNums = strings.Split(nums, ",")
	}

	size := len(sNums)
	nodes := make([]int, size)
	for i, num := range sNums {
		num = strings.TrimSpace(num)
		atoi, err := strconv.Atoi(num)
		throughErrorPanic(err)
		nodes[i] = atoi
	}
	return nodes
}

func strToByteArray(str string) []byte {
	str = fixBrackets(str)
	if str == "" {
		return nil
	}

	sb := strings.Split(str, ",")
	size := len(sb)
	ret := make([]byte, size)
	for i, b := range sb {
		fmt.Print(b)
		if b == "" {
			throughErrorPanic(errors.New("found null characters"))
		}
		ret[i] = b[0]
	}
	return ret
}

func strToIntArrArray(nums string) [][]int {
	nums = fixBrackets(nums)
	if nums == "" {
		return nil
	}

	var ret [][]int
	sta := strings.Index(nums, "[")
	end := strings.Index(nums, "]")
	for end != -1 && sta != -1 {
		ret = append(ret, strToIntArray(nums[sta:end]))
		nums = nums[end+1:]
		sta = strings.Index(nums, "[")
		end = strings.Index(nums, "]")
	}

	return ret
}

func strToListNode(nums string, pos int) *ListNode {
	nums = fixBrackets(nums)
	var sNums []string
	if nums == "" {
		sNums = make([]string, 1)
	} else {
		if strings.Contains(nums, "->") {
			sNums = strings.Split(nums, "->")
		} else {
			sNums = strings.Split(nums, ",")
		}
	}
	size := len(sNums)
	nodes := make([]ListNode, size)
	for i := range nodes {
		if "null" == strings.ToLower(sNums[i]) {
			nodes = nodes[:i-1]
			break
		}
		if sNums[i] == "" {
			continue
		}
		n, err := strconv.Atoi(sNums[i])
		throughErrorPanic(err)
		nodes[i].Val = n
		//fmt.Println("nodes", i, nodes[i])
		if i < size-1 {
			nodes[i].Next = &nodes[i+1]
		} else {
			if pos >= 0 && pos < size {
				nodes[i].Next = &nodes[pos]
				//fmt.Println("nodes pos", pos, nodes[pos])
			}
		}
	}
	return &nodes[0]
}

func strToTreeNode(nums string) *TreeNode {
	nums = fixBrackets(nums)
	var sNums []string
	if nums == "" {
		sNums = make([]string, 1)
	} else {
		sNums = strings.Split(nums, ",")
	}

	size := len(sNums)
	nodes := make([]*TreeNode, size)
	for idx, num := range sNums {
		if num == strings.ToLower("null") || num == "" {
			continue
		}

		i, err := strconv.Atoi(num)
		throughErrorPanic(err)
		nodes[idx] = &TreeNode{Val: i}
	}

	i := 1
	queue := make([]*TreeNode, 1, size*2)
	queue[0] = nodes[0]
	for i < len(sNums) {
		node := queue[0]
		queue = queue[1:]

		if i < len(sNums) && nodes[i] != nil {
			node.Left = nodes[i]
			queue = append(queue, node.Left)
		}
		i++

		if i < len(sNums) && nodes[i] != nil {
			node.Right = nodes[i]
			queue = append(queue, node.Right)
		}
		i++
	}
	return nodes[0]
}

func throughErrorPanic(err error) {
	if err != nil {
		panic(err)
	}
}
