package _2810

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_finalString(t *testing.T) {
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
			if got := finalString(tt.args.s); got != tt.want {
				t.Errorf("finalString() = %v, want %v", got, tt.want)
			}
		})
	}
}
