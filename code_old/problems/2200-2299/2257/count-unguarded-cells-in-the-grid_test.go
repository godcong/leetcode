package _2257

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_countUnguarded(t *testing.T) {
	type args struct {
		m      int
		n      int
		guards [][]int
		walls  [][]int
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
			if got := countUnguarded(tt.args.m, tt.args.n, tt.args.guards, tt.args.walls); got != tt.want {
				t.Errorf("countUnguarded() = %v, want %v", got, tt.want)
			}
		})
	}
}
