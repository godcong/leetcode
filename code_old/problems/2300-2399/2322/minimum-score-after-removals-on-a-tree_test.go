package _2322

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_minimumScore(t *testing.T) {
	type args struct {
		nums  []int
		edges [][]int
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
			if got := minimumScore(tt.args.nums, tt.args.edges); got != tt.want {
				t.Errorf("minimumScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
