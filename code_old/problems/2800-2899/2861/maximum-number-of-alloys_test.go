package _2861

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_maxNumberOfAlloys(t *testing.T) {
	type args struct {
		n           int
		k           int
		budget      int
		composition [][]int
		stock       []int
		cost        []int
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
			if got := maxNumberOfAlloys(tt.args.n, tt.args.k, tt.args.budget, tt.args.composition, tt.args.stock, tt.args.cost); got != tt.want {
				t.Errorf("maxNumberOfAlloys() = %v, want %v", got, tt.want)
			}
		})
	}
}
