package _1458

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_maxDotProduct(t *testing.T) {
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
			if got := maxDotProduct(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("maxDotProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
