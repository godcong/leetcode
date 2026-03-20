package _0846

import (
	"sort"
)

func isNStraightHand(hand []int, groupSize int) bool {
	if len(hand)%groupSize != 0 {
		return false
	}
	if groupSize == 1 {
		return true
	}

	sort.Ints(hand)
	for i := 0; i < len(hand); i++ {
		if hand[i] == -1 {
			continue
		}

		value, cnt, isFull := hand[i], 0, false
		for j := i; j < len(hand); j++ {
			if hand[j] == -1 {
				continue
			}

			if hand[j] == value+cnt {
				hand[j] = -1
				cnt++
				if cnt == groupSize {
					isFull = true
					break
				}
			}

			if hand[j] >= value+groupSize {
				return false
			}
		}

		if !isFull {
			return false
		}
	}
	return true
}
