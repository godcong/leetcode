package _1163

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_lastSubstring(t *testing.T) {
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
			if got := lastSubstring(tt.args.s); got != tt.want {
				t.Errorf("lastSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
