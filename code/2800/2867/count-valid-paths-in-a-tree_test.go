package _2867

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_countPaths(t *testing.T) {
	type args struct {
		n     int
		edges [][]int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPaths(tt.args.n, tt.args.edges); got != tt.want {
				t.Errorf("countPaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
