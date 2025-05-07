package _3341

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_minTimeToReach(t *testing.T) {
	type args struct {
		moveTime [][]int
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
			if got := minTimeToReach(tt.args.moveTime); got != tt.want {
				t.Errorf("minTimeToReach() = %v, want %v", got, tt.want)
			}
		})
	}
}
