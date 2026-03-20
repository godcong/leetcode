package _2734

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_smallestString(t *testing.T) {
	type args struct {
		s string
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
			if got := smallestString(tt.args.s); got != tt.want {
				t.Errorf("smallestString() = %v, want %v", got, tt.want)
			}
		})
	}
}
