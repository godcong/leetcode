package _3239

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_minFlips(t *testing.T) {
	type args struct {
		grid [][]int
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
			if got := minFlips(tt.args.grid); got != tt.want {
				t.Errorf("minFlips() = %v, want %v", got, tt.want)
			}
		})
	}
}
