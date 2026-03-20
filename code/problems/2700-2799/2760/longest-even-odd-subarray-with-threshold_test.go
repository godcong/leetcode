package _2760

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_longestAlternatingSubarray(t *testing.T) {
	type args struct {
		nums      []int
		threshold int
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
			if got := longestAlternatingSubarray(tt.args.nums, tt.args.threshold); got != tt.want {
				t.Errorf("longestAlternatingSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}
