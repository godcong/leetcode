package _3218

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_minimumCost(t *testing.T) {
	type args struct {
		m             int
		n             int
		horizontalCut []int
		verticalCut   []int
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
			if got := minimumCost(tt.args.m, tt.args.n, tt.args.horizontalCut, tt.args.verticalCut); got != tt.want {
				t.Errorf("minimumCost() = %v, want %v", got, tt.want)
			}
		})
	}
}
