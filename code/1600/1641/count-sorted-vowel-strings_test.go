package _1641

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_countVowelStrings(t *testing.T) {
	type args struct {
		n int
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
			if got := countVowelStrings(tt.args.n); got != tt.want {
				t.Errorf("countVowelStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
