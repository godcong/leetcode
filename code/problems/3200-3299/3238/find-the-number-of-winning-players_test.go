package _3238

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_winningPlayerCount(t *testing.T) {
	type args struct {
		n    int
		pick [][]int
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
			if got := winningPlayerCount(tt.args.n, tt.args.pick); got != tt.want {
				t.Errorf("winningPlayerCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
