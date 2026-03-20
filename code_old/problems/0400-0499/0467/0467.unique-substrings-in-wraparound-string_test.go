package _0467

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_findSubstringInWraproundString(t *testing.T) {
	type args struct {
		p string
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
			if got := findSubstringInWraproundString(tt.args.p); got != tt.want {
				t.Errorf("findSubstringInWraproundString() = %v, want %v", got, tt.want)
			}
		})
	}
}
