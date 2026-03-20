package _1654

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_minimumJumps(t *testing.T) {
	type args struct {
		forbidden []int
		a         int
		b         int
		x         int
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
			if got := minimumJumps(tt.args.forbidden, tt.args.a, tt.args.b, tt.args.x); got != tt.want {
				t.Errorf("minimumJumps() = %v, want %v", got, tt.want)
			}
		})
	}
}
