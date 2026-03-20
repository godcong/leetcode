package _0675

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_cutOffTree(t *testing.T) {
	type args struct {
		forest [][]int
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
			if got := cutOffTree(tt.args.forest); got != tt.want {
				t.Errorf("cutOffTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
