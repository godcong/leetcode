package _3356

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_minZeroArray(t *testing.T) {
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
			if got := minZeroArray(tt.args.nums, tt.args.queries); got != tt.want {
				t.Errorf("minZeroArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
