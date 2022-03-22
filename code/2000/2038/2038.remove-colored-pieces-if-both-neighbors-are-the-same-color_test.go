package _2038

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_winnerOfGame(t *testing.T) {
	type args struct {
		colors string
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
			if got := winnerOfGame(tt.args.colors); got != tt.want {
				t.Errorf("winnerOfGame() = %v, want %v", got, tt.want)
			}
		})
	}
}
