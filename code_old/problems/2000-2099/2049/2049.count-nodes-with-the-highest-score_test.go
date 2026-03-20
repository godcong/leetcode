package _2049

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_countHighestScoreNodes(t *testing.T) {
	type args struct {
		parents []int
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
			if got := countHighestScoreNodes(tt.args.parents); got != tt.want {
				t.Errorf("countHighestScoreNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
