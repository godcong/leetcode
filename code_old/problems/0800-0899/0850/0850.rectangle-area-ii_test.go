package _0850

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_rectangleArea(t *testing.T) {
	type args struct {
		rectangles [][]int
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
			if got := rectangleArea(tt.args.rectangles); got != tt.want {
				t.Errorf("rectangleArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
