package _3714

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_longestBalanced(t *testing.T) {
	type args struct {
		s string
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
			if got := longestBalanced(tt.args.s); got != tt.want {
				t.Errorf("longestBalanced() = %v, want %v", got, tt.want)
			}
		})
	}
}
