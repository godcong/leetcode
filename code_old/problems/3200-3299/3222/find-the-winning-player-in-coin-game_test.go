package _3222

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_losingPlayer(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := losingPlayer(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("losingPlayer() = %v, want %v", got, tt.want)
			}
		})
	}
}
