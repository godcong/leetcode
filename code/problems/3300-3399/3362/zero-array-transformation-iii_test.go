package _3362

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_maxRemoval(t *testing.T) {
	type args struct {
		nums    []int
		queries [][]int
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
			if got := maxRemoval(tt.args.nums, tt.args.queries); got != tt.want {
				t.Errorf("maxRemoval() = %v, want %v", got, tt.want)
			}
		})
	}
}
