package _3011

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_canSortArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canSortArray(tt.args.nums); got != tt.want {
				t.Errorf("canSortArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
