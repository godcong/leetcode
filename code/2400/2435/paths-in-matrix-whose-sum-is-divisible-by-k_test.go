package _2435

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_numberOfPaths(t *testing.T) {
	type args struct {
		grid [][]int
		k    int
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
			if got := numberOfPaths(tt.args.grid, tt.args.k); got != tt.want {
				t.Errorf("numberOfPaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
