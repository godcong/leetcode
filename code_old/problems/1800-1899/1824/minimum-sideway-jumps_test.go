package _1824

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_minSideJumps(t *testing.T) {
	type args struct {
		obstacles []int
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
			if got := minSideJumps(tt.args.obstacles); got != tt.want {
				t.Errorf("minSideJumps() = %v, want %v", got, tt.want)
			}
		})
	}
}
