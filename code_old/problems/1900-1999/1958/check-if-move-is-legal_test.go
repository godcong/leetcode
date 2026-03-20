package _1958

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_checkMove(t *testing.T) {
	type args struct {
		board [][]byte
		rMove int
		cMove int
		color byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkMove(tt.args.board, tt.args.rMove, tt.args.cMove, tt.args.color); got != tt.want {
				t.Errorf("checkMove() = %v, want %v", got, tt.want)
			}
		})
	}
}
