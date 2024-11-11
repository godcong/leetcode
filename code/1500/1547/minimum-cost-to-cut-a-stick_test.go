package _1547

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_minCost(t *testing.T) {
	type args struct {
		n    int
		cuts []int
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
			if got := minCost(tt.args.n, tt.args.cuts); got != tt.want {
				t.Errorf("minCost() = %v, want %v", got, tt.want)
			}
		})
	}
}
