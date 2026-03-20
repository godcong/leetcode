package _1572

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_diagonalSum(t *testing.T) {
	type args struct {
		mat [][]int
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
			if got := diagonalSum(tt.args.mat); got != tt.want {
				t.Errorf("diagonalSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
