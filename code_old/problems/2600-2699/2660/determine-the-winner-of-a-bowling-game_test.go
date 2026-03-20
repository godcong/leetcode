package _2660

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_isWinner(t *testing.T) {
	type args struct {
		player1 []int
		player2 []int
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
			if got := isWinner(tt.args.player1, tt.args.player2); got != tt.want {
				t.Errorf("isWinner() = %v, want %v", got, tt.want)
			}
		})
	}
}
