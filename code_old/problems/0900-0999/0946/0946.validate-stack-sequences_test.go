package _0946

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_validateStackSequences(t *testing.T) {
	type args struct {
		pushed []int
		popped []int
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
			if got := validateStackSequences(tt.args.pushed, tt.args.popped); got != tt.want {
				t.Errorf("validateStackSequences() = %v, want %v", got, tt.want)
			}
		})
	}
}
