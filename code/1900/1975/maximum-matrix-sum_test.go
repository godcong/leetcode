package _1975

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_maxMatrixSum(t *testing.T) {
	type args struct {
		matrix [][]int
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
			if got := maxMatrixSum(tt.args.matrix); got != tt.want {
				t.Errorf("maxMatrixSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
