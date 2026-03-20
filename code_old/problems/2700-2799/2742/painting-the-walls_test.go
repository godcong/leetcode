package _2742

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_paintWalls(t *testing.T) {
	type args struct {
		cost []int
		time []int
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
			if got := paintWalls(tt.args.cost, tt.args.time); got != tt.want {
				t.Errorf("paintWalls() = %v, want %v", got, tt.want)
			}
		})
	}
}
