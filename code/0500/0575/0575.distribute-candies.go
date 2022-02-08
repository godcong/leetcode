package _0575

func distributeCandies(candyType []int) int {
	max := len(candyType) / 2
	set := make(map[int]bool, len(candyType)/2)

	for _, t := range candyType {
		set[t] = true
		if len(set) >= max {
			return max
		}
	}
	return len(set)
}
