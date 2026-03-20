package _0882

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_reachableNodes(t *testing.T) {
	type args struct {
		edges    [][]int
		maxMoves int
		n        int
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
			if got := reachableNodes(tt.args.edges, tt.args.maxMoves, tt.args.n); got != tt.want {
				t.Errorf("reachableNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
