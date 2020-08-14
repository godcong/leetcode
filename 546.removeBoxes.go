package leetcode

/*
546. 移除盒子
给出一些不同颜色的盒子，盒子的颜色由数字表示，即不同的数字表示不同的颜色。
你将经过若干轮操作去去掉盒子，直到所有的盒子都去掉为止。每一轮你可以移除具有相同颜色的连续 k 个盒子（k >= 1），这样一轮之后你将得到 k*k 个积分。
当你将所有盒子都去掉之后，求你能获得的最大积分和。



示例：

输入：boxes = [1,3,2,2,2,3,4,3,1]
输出：23
解释：
[1, 3, 2, 2, 2, 3, 4, 3, 1]
----> [1, 3, 3, 4, 3, 1] (3*3=9 分)
----> [1, 3, 3, 3, 1] (1*1=1 分)
----> [1, 1] (3*3=9 分)
----> [] (2*2=4 分)


提示：

1 <= boxes.length <= 100
1 <= boxes[i] <= 100
*/
func removeBoxes(boxes []int) int {
	//size := len(boxes)
	//if size < 2 {
	//	return 1
	//}
	//
	//tmpBox := make([]int, len(boxes))
	//last := boxes[0]
	//count, max := 1, 0
	//
	//for i := 1; i < size; i++ {
	//	if boxes[i] != last {
	//		copy(tmpBox, boxes[:i-count])
	//		copy(tmpBox[i-count:], boxes[i:])
	//		if v := removeBoxesAction(tmpBox[:size-count]); (v + count*count) > max {
	//			max = v + count*count
	//		}
	//		count, last = 1, boxes[i]
	//		continue
	//	}
	//	count, last = count+1, boxes[i]
	//}
	//if count > 1 {
	//	return max + count*count
	//}
	//
	//return max
	return removeBoxesAction(boxes)
}

func removeBoxesAction(boxes []int) int {
	size := len(boxes)
	if size < 2 {
		return 1
	}

	tmpBox := make([]int, len(boxes))
	last := boxes[0]
	count, max, maxCount := 1, 0, false

	for i := 1; i < size; i++ {
		if boxes[i] != last {
			copy(tmpBox, boxes[:i-count])
			copy(tmpBox[i-count:], boxes[i:])
			if v := removeBoxesAction(tmpBox[:size-count]); (v + count*count) > max {
				max = v + count*count
				maxCount = true
			}
			count, last = 1, boxes[i]
			continue
		}
		count, last = count+1, boxes[i]
	}
	if !maxCount {
		return max + count*count
	}
	return max
}
