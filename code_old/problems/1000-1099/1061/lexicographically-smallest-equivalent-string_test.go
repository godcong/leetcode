package _1061

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_smallestEquivalentString(t *testing.T) {
	type args struct {
		s1      string
		s2      string
		baseStr string
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
			if got := smallestEquivalentString(tt.args.s1, tt.args.s2, tt.args.baseStr); got != tt.want {
				t.Errorf("smallestEquivalentString() = %v, want %v", got, tt.want)
			}
		})
	}
}
