package _1996

func numberOfWeakCharacters(properties [][]int) int {
	maxAttack := 0
	for _, x := range properties {
		if maxAttack < x[0] {
			maxAttack = x[0]
		}
	}

	attacks := make([]int, maxAttack+1)
	for _, p := range properties {
		if attacks[p[0]] < p[1] {
			attacks[p[0]] = p[1]
		}
	}

	maxDefence := 0
	for i := maxAttack; i >= 0; i-- {
		temp := maxDefence
		if attacks[i] > maxDefence {
			maxDefence = attacks[i]
		}
		attacks[i] = temp
	}
	ans := 0
	for _, p := range properties {
		if p[0] < maxAttack && p[1] < attacks[p[0]] {
			ans++
		}
	}
	return ans
}
