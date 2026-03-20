package _2609

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_findTheLongestBalancedSubstring(t *testing.T) {
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
			if got := findTheLongestBalancedSubstring(tt.args.s); got != tt.want {
				t.Errorf("findTheLongestBalancedSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
