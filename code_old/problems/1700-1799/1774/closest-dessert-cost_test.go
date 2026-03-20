package _1774

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_closestCost(t *testing.T) {
	type args struct {
		baseCosts    []int
		toppingCosts []int
		target       int
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
			if got := closestCost(tt.args.baseCosts, tt.args.toppingCosts, tt.args.target); got != tt.want {
				t.Errorf("closestCost() = %v, want %v", got, tt.want)
			}
		})
	}
}
