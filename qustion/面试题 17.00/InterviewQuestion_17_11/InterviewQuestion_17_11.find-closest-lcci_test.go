package _InterviewQuestion_17_11

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_findClosest(t *testing.T) {
	type args struct {
		words []string
		word1 string
		word2 string
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
			if got := findClosest(tt.args.words, tt.args.word1, tt.args.word2); got != tt.want {
				t.Errorf("findClosest() = %v, want %v", got, tt.want)
			}
		})
	}
}
