package _0551

func checkRecord(s string) bool {
	absents, lates := 0, 0
	for _, ch := range s {
		if ch == 'A' {
			absents++
			if absents >= 2 {
				return false
			}
		}
		if ch == 'L' {
			lates++
			if lates >= 3 {
				return false
			}
		} else {
			lates = 0
		}
	}
	return true
}
