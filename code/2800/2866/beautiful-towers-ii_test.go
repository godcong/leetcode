package _2866

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_maximumSumOfHeights(t *testing.T) {
	type args struct {
		maxHeights []int
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
			if got := maximumSumOfHeights(tt.args.maxHeights); got != tt.want {
				t.Errorf("maximumSumOfHeights() = %v, want %v", got, tt.want)
			}
		})
	}
}
