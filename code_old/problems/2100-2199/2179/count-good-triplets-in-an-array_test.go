package _2179

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_goodTriplets(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
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
			if got := goodTriplets(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("goodTriplets() = %v, want %v", got, tt.want)
			}
		})
	}
}
