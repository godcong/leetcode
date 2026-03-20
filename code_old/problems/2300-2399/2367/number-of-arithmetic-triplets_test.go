package _2367

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_arithmeticTriplets(t *testing.T) {
	type args struct {
		nums []int
		diff int
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
			if got := arithmeticTriplets(tt.args.nums, tt.args.diff); got != tt.want {
				t.Errorf("arithmeticTriplets() = %v, want %v", got, tt.want)
			}
		})
	}
}
