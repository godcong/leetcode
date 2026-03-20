package _0591

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_isValid(t *testing.T) {
	type args struct {
		code string
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
			if got := isValid(tt.args.code); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
