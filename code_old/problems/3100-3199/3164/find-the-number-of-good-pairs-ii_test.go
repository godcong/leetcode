package _3164

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_numberOfPairs(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
		k     int
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
			if got := numberOfPairs(tt.args.nums1, tt.args.nums2, tt.args.k); got != tt.want {
				t.Errorf("numberOfPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
