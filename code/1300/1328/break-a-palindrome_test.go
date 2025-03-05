package _1328

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_breakPalindrome(t *testing.T) {
	type args struct {
		palindrome string
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
			if got := breakPalindrome(tt.args.palindrome); got != tt.want {
				t.Errorf("breakPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
