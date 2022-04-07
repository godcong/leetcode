package _0796

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_rotateString(t *testing.T) {
	type args struct {
		s    string
		goal string
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
			if got := rotateString(tt.args.s, tt.args.goal); got != tt.want {
				t.Errorf("rotateString() = %v, want %v", got, tt.want)
			}
		})
	}
}
