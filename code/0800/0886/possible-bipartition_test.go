package _0886

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_possibleBipartition(t *testing.T) {
	type args struct {
		n        int
		dislikes [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := possibleBipartition(tt.args.n, tt.args.dislikes); got != tt.want {
				t.Errorf("possibleBipartition() = %v, want %v", got, tt.want)
			}
		})
	}
}
