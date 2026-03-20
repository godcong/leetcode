package _0410

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_splitArray(t *testing.T) {
	type args struct {
		nums []int
		k    int
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
			if got := splitArray(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("splitArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
