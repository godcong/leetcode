package _2845

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_countInterestingSubarrays(t *testing.T) {
	type args struct {
		nums   []int
		modulo int
		k      int
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
			if got := countInterestingSubarrays(tt.args.nums, tt.args.modulo, tt.args.k); got != tt.want {
				t.Errorf("countInterestingSubarrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
