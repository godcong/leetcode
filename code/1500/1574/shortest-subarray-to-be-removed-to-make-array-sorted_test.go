package _1574

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_findLengthOfShortestSubarray(t *testing.T) {
	type args struct {
		arr []int
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
			if got := findLengthOfShortestSubarray(tt.args.arr); got != tt.want {
				t.Errorf("findLengthOfShortestSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}
