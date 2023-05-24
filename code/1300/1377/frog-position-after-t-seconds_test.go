package _1377

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_frogPosition(t *testing.T) {
	type args struct {
		n      int
		edges  [][]int
		t      int
		target int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := frogPosition(tt.args.n, tt.args.edges, tt.args.t, tt.args.target); got != tt.want {
				t.Errorf("frogPosition() = %v, want %v", got, tt.want)
			}
		})
	}
}
