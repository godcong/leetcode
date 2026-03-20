package _3146

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_findPermutationDifference(t *testing.T) {
	type args struct {
		s string
		t string
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
			if got := findPermutationDifference(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("findPermutationDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
