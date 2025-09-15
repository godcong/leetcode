package _1935

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_canBeTypedWords(t *testing.T) {
	type args struct {
		text          string
		brokenLetters string
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
			if got := canBeTypedWords(tt.args.text, tt.args.brokenLetters); got != tt.want {
				t.Errorf("canBeTypedWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
