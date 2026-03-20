package _0833

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_findReplaceString(t *testing.T) {
	type args struct {
		s       string
		indices []int
		sources []string
		targets []string
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
			if got := findReplaceString(tt.args.s, tt.args.indices, tt.args.sources, tt.args.targets); got != tt.want {
				t.Errorf("findReplaceString() = %v, want %v", got, tt.want)
			}
		})
	}
}
