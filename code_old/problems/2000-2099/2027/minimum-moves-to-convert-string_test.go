package _2027

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_minimumMoves(t *testing.T) {
	type args struct {
		s string
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
			if got := minimumMoves(tt.args.s); got != tt.want {
				t.Errorf("minimumMoves() = %v, want %v", got, tt.want)
			}
		})
	}
}
