package _1138

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_alphabetBoardPath(t *testing.T) {
	type args struct {
		target string
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
			if got := alphabetBoardPath(tt.args.target); got != tt.want {
				t.Errorf("alphabetBoardPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
