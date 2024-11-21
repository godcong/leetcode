package _3248

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_finalPositionOfSnake(t *testing.T) {
	type args struct {
		n        int
		commands []string
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
			if got := finalPositionOfSnake(tt.args.n, tt.args.commands); got != tt.want {
				t.Errorf("finalPositionOfSnake() = %v, want %v", got, tt.want)
			}
		})
	}
}
