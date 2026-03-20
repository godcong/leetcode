package _2581

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_rootCount(t *testing.T) {
	type args struct {
		edges   [][]int
		guesses [][]int
		k       int
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
			if got := rootCount(tt.args.edges, tt.args.guesses, tt.args.k); got != tt.want {
				t.Errorf("rootCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
