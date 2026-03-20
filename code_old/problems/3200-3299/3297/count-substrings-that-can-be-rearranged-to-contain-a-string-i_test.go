package _3297

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_validSubstringCount(t *testing.T) {
	type args struct {
		word1 string
		word2 string
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
			if got := validSubstringCount(tt.args.word1, tt.args.word2); got != tt.want {
				t.Errorf("validSubstringCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
