package _2707

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_minExtraChar(t *testing.T) {
	type args struct {
		s          string
		dictionary []string
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
			if got := minExtraChar(tt.args.s, tt.args.dictionary); got != tt.want {
				t.Errorf("minExtraChar() = %v, want %v", got, tt.want)
			}
		})
	}
}
