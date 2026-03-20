package _2207

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_maximumSubsequenceCount(t *testing.T) {
	type args struct {
		text    string
		pattern string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumSubsequenceCount(tt.args.text, tt.args.pattern); got != tt.want {
				t.Errorf("maximumSubsequenceCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
