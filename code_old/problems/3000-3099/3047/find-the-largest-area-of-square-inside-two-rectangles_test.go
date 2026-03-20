package _3047

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_largestSquareArea(t *testing.T) {
	type args struct {
		bottomLeft [][]int
		topRight   [][]int
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
			if got := largestSquareArea(tt.args.bottomLeft, tt.args.topRight); got != tt.want {
				t.Errorf("largestSquareArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
