package _2154

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_findFinalValue(t *testing.T) {
	type args struct {
		nums     []int
		original int
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
			if got := findFinalValue(tt.args.nums, tt.args.original); got != tt.want {
				t.Errorf("findFinalValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
