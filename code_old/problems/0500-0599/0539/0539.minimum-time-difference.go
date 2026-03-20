package _0539

func findMinDifference(timePoints []string) int {
	n := len(timePoints) - 1
	if n > totalMinutes {
		return 0
	}
	tmp := make([]int, totalMinutes)
	for _, tp := range timePoints {
		t := minutes(tp)
		if tmp[t] != 0 {
			return 0
		}
		tmp[t] = 1
	}
	result := totalMinutes + 1
	first, last := -1, -1
	for j, v := range tmp {
		if v == 0 {
			continue
		}
		if first == -1 {
			first = j
		}
		if last != -1 {
			result = min(result, j-last)
		}
		last = j
	}
	result = min(result, first+totalMinutes-last)
	return result
}

const totalMinutes = 24 * 60

func minutes(s string) int {
	return int(s[0]-'0')*600 + int(s[1]-'0')*60 + int(s[3]-'0')*10 + int(s[4]-'0')
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
