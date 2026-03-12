package _3600

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_maxStability(t *testing.T) {
	type args struct {
		n     int
		edges [][]int
		k     int
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
			if got := maxStability(tt.args.n, tt.args.edges, tt.args.k); got != tt.want {
				t.Errorf("maxStability() = %v, want %v", got, tt.want)
			}
		})
	}
}
