package _2787

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_numberOfWays(t *testing.T) {
	type args struct {
		n int
		x int
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
			if got := numberOfWays(tt.args.n, tt.args.x); got != tt.want {
				t.Errorf("numberOfWays() = %v, want %v", got, tt.want)
			}
		})
	}
}
