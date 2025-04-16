package _2537

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_countGood(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countGood(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("countGood() = %v, want %v", got, tt.want)
			}
		})
	}
}
