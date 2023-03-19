package _1625

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_findLexSmallestString(t *testing.T) {
	type args struct {
		s string
		a int
		b int
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
			if got := findLexSmallestString(tt.args.s, tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("findLexSmallestString() = %v, want %v", got, tt.want)
			}
		})
	}
}
