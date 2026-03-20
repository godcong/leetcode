package _0756

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_pyramidTransition(t *testing.T) {
	type args struct {
		bottom  string
		allowed []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pyramidTransition(tt.args.bottom, tt.args.allowed); got != tt.want {
				t.Errorf("pyramidTransition() = %v, want %v", got, tt.want)
			}
		})
	}
}
