package _2586

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_vowelStrings(t *testing.T) {
	type args struct {
		words []string
		left  int
		right int
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
			if got := vowelStrings(tt.args.words, tt.args.left, tt.args.right); got != tt.want {
				t.Errorf("vowelStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
