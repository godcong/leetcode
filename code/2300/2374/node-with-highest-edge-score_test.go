package _2374

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_edgeScore(t *testing.T) {
	type args struct {
		edges []int
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
			if got := edgeScore(tt.args.edges); got != tt.want {
				t.Errorf("edgeScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
