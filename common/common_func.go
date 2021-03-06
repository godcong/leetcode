package common

import (
	"errors"
	"strconv"
	"strings"
)

func FixBrackets(nums string) string {
	if nums == "" {
		return ""
	}
	nums = strings.TrimSpace(nums)
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

func StrToIntArray(nums string) []int {
	nums = FixBrackets(nums)
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
		if num == "null" {
			nodes[i] = 0
			continue
		}
		atoi, err := strconv.Atoi(num)
		throughErrorPanic(err)
		nodes[i] = atoi
	}
	return nodes
}

func StrToByteArrayArray(str string) [][]byte {
	str = FixBrackets(str)
	if str == "" {
		return nil
	}

	var ret [][]byte
	sta := strings.Index(str, "[")
	end := strings.Index(str, "]")
	for end != -1 && sta != -1 {
		ret = append(ret, StrToByteArray(str[sta:end]))
		str = str[end+1:]
		sta = strings.Index(str, "[")
		end = strings.Index(str, "]")
	}

	return ret
}

func StrToByteArray(str string) []byte {
	str = FixBrackets(str)
	if str == "" {
		return nil
	}

	sb := strings.Split(str, ",")
	size := len(sb)
	ret := make([]byte, size)
	for i, b := range sb {
		if b == "" {
			throughErrorPanic(errors.New("found null characters"))
		}
		ret[i] = b[1]
	}
	return ret
}

func StrToStringArray(str string) []string {
	str = FixBrackets(str)
	if str == "" {
		return nil
	}
	str = strings.ReplaceAll(str, "\"", "")
	str = strings.Trim(str, "\n")
	ret := strings.Split(str, ",")
	for i, s := range ret {
		ret[i] = strings.TrimSpace(s)
	}
	return ret
}

func StrToStringArrayArray(str string) [][]string {
	str = FixBrackets(str)
	sta := strings.Index(str, "[")
	end := strings.Index(str, "]")
	var ret [][]string
	for sta != -1 && end != -1 {
		ret = append(ret, StrToStringArray(str[sta:end+1]))
		str = str[end+1:]
		sta = strings.Index(str, "[")
		end = strings.Index(str, "]")
	}
	return ret
}

func StrToIntArrArray(nums string) [][]int {
	nums = FixBrackets(nums)
	if nums == "" {
		return nil
	}

	var ret [][]int
	sta := strings.Index(nums, "[")
	end := strings.Index(nums, "]")
	for end != -1 && sta != -1 {
		ret = append(ret, StrToIntArray(nums[sta:end]))
		nums = nums[end+1:]
		sta = strings.Index(nums, "[")
		end = strings.Index(nums, "]")
	}

	return ret
}

func StrToListNode(nums string, pos int) *ListNode {
	nums = FixBrackets(nums)
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

func StrToTreeNode(nums string) *TreeNode {
	nums = FixBrackets(nums)
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

func StrToNestedIntegerList(str string) []*NestedInteger {
	str = FixBrackets(str)
	var sNums []string
	if str == "" {
		sNums = make([]string, 1)
	} else {
		sNums = strings.Split(str, ",")
	}

	size := len(sNums)
	nodes := make([]*NestedInteger, size)
	//todo

	return nodes
}

func throughErrorPanic(err error) {
	if err != nil {
		panic(err)
	}
}

func StrToEmployees(str string) []*Employee {
	str = FixBrackets(str)
	return nil
}
