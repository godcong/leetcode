package _3127

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_canMakeSquare(t *testing.T) {
	type args struct {
		grid [][]byte
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
			if got := canMakeSquare(tt.args.grid); got != tt.want {
				t.Errorf("canMakeSquare() = %v, want %v", got, tt.want)
			}
		})
	}
}
