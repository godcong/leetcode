package _3000

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_areaOfMaxDiagonal(t *testing.T) {
	type args struct {
		dimensions [][]int
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
			if got := areaOfMaxDiagonal(tt.args.dimensions); got != tt.want {
				t.Errorf("areaOfMaxDiagonal() = %v, want %v", got, tt.want)
			}
		})
	}
}
