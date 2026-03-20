package _3397

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_maxDistinctElements(t *testing.T) {
	type args struct {
		nums []int
		k    int
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
			if got := maxDistinctElements(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("maxDistinctElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
