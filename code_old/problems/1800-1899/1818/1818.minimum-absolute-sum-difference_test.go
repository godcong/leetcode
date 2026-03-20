package _1818

import (
	"testing"
)

func Test_minAbsoluteSumDiff(t *testing.T) {
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
			if got := minAbsoluteSumDiff(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("minAbsoluteSumDiff() = %v, want %v", got, tt.want)
			}
		})
	}
}
