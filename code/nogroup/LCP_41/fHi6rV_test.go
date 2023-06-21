package _LCP_41

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_flipChess(t *testing.T) {
	type args struct {
		chessboard []string
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
			if got := flipChess(tt.args.chessboard); got != tt.want {
				t.Errorf("flipChess() = %v, want %v", got, tt.want)
			}
		})
	}
}
