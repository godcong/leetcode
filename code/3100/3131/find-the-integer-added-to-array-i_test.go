package _3131

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_addedInteger(t *testing.T) {
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
			if got := addedInteger(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("addedInteger() = %v, want %v", got, tt.want)
			}
		})
	}
}
