package _2605

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_minNumber(t *testing.T) {
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
			if got := minNumber(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("minNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
