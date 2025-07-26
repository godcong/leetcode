package _3480

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_maxSubarrays(t *testing.T) {
	type args struct {
		n                int
		conflictingPairs [][]int
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
			if got := maxSubarrays(tt.args.n, tt.args.conflictingPairs); got != tt.want {
				t.Errorf("maxSubarrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
