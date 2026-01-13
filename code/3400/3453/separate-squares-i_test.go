package _3453

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_separateSquares(t *testing.T) {
	type args struct {
		squares [][]int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := separateSquares(tt.args.squares); got != tt.want {
				t.Errorf("separateSquares() = %v, want %v", got, tt.want)
			}
		})
	}
}
