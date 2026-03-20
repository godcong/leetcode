package _1466

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_minReorder(t *testing.T) {
	type args struct {
		n           int
		connections [][]int
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
			if got := minReorder(tt.args.n, tt.args.connections); got != tt.want {
				t.Errorf("minReorder() = %v, want %v", got, tt.want)
			}
		})
	}
}
