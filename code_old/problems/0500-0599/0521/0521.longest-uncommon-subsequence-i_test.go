package _0521

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_findLUSlength(t *testing.T) {
	type args struct {
		a string
		b string
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
			if got := findLUSlength(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("findLUSlength() = %v, want %v", got, tt.want)
			}
		})
	}
}
