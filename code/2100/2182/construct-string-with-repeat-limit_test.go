package _2182

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_repeatLimitedString(t *testing.T) {
	type args struct {
		s           string
		repeatLimit int
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
			if got := repeatLimitedString(tt.args.s, tt.args.repeatLimit); got != tt.want {
				t.Errorf("repeatLimitedString() = %v, want %v", got, tt.want)
			}
		})
	}
}
