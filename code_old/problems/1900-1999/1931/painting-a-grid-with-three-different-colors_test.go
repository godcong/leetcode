package _1931

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_colorTheGrid(t *testing.T) {
	type args struct {
		m int
		n int
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
			if got := colorTheGrid(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("colorTheGrid() = %v, want %v", got, tt.want)
			}
		})
	}
}
