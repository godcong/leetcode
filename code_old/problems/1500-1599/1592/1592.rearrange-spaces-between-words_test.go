package _1592

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_reorderSpaces(t *testing.T) {
	type args struct {
		text string
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
			if got := reorderSpaces(tt.args.text); got != tt.want {
				t.Errorf("reorderSpaces() = %v, want %v", got, tt.want)
			}
		})
	}
}
