package _1719

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_checkWays(t *testing.T) {
	type args struct {
		pairs [][]int
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
			if got := checkWays(tt.args.pairs); got != tt.want {
				t.Errorf("checkWays() = %v, want %v", got, tt.want)
			}
		})
	}
}
