package _3562

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_maxProfit(t *testing.T) {
	type args struct {
		n         int
		present   []int
		future    []int
		hierarchy [][]int
		budget    int
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
			if got := maxProfit(tt.args.n, tt.args.present, tt.args.future, tt.args.hierarchy, tt.args.budget); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
