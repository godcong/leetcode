package _0782

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_movesToChessboard(t *testing.T) {
	type args struct {
		board [][]int
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
			if got := movesToChessboard(tt.args.board); got != tt.want {
				t.Errorf("movesToChessboard() = %v, want %v", got, tt.want)
			}
		})
	}
}
