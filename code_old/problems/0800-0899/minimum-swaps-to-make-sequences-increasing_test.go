package _0801

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_minSwap(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
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
			if got := minSwap(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("minSwap() = %v, want %v", got, tt.want)
			}
		})
	}
}
