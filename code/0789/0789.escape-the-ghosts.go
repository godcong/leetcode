package _0789

func escapeGhosts(ghosts [][]int, target []int) bool {
	source := []int{0, 0}
	distance := calDis(source, target)
	for _, item := range ghosts {
		ghostDis := calDis(item, target)
		if ghostDis <= distance {
			return false
		}
	}
	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func calDis(source, target []int) int {
	return abs(target[0]-source[0]) + abs(target[1]-source[1])
}
