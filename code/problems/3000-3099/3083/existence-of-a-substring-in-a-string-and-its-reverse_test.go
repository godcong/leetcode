package _3083

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_isSubstringPresent(t *testing.T) {
	type args struct {
		s string
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
			if got := isSubstringPresent(tt.args.s); got != tt.want {
				t.Errorf("isSubstringPresent() = %v, want %v", got, tt.want)
			}
		})
	}
}
