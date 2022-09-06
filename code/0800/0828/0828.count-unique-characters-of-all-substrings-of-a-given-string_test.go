package _0828

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_uniqueLetterString(t *testing.T) {
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
			if got := uniqueLetterString(tt.args.s); got != tt.want {
				t.Errorf("uniqueLetterString() = %v, want %v", got, tt.want)
			}
		})
	}
}
