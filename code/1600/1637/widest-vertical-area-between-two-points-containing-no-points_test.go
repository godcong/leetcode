package _1637

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_maxWidthOfVerticalArea(t *testing.T) {
	type args struct {
		points [][]int
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
			if got := maxWidthOfVerticalArea(tt.args.points); got != tt.want {
				t.Errorf("maxWidthOfVerticalArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
