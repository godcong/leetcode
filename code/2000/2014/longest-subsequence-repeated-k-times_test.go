package _2014

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_longestSubsequenceRepeatedK(t *testing.T) {
	type args struct {
		s string
		k int
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
			if got := longestSubsequenceRepeatedK(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("longestSubsequenceRepeatedK() = %v, want %v", got, tt.want)
			}
		})
	}
}
