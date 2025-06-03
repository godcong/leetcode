package _1298

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_maxCandies(t *testing.T) {
	type args struct {
		status         []int
		candies        []int
		keys           [][]int
		containedBoxes [][]int
		initialBoxes   []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxCandies(tt.args.status, tt.args.candies, tt.args.keys, tt.args.containedBoxes, tt.args.initialBoxes); got != tt.want {
				t.Errorf("maxCandies() = %v, want %v", got, tt.want)
			}
		})
	}
}
