package _2040

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_kthSmallestProduct(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
		k     int64
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
			if got := kthSmallestProduct(tt.args.nums1, tt.args.nums2, tt.args.k); got != tt.want {
				t.Errorf("kthSmallestProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
