package _2368

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_reachableNodes(t *testing.T) {
	type args struct {
		n          int
		edges      [][]int
		restricted []int
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
			if got := reachableNodes(tt.args.n, tt.args.edges, tt.args.restricted); got != tt.want {
				t.Errorf("reachableNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
