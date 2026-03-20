package _3249

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_countGoodNodes(t *testing.T) {
	type args struct {
		edges [][]int
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
			if got := countGoodNodes(tt.args.edges); got != tt.want {
				t.Errorf("countGoodNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
