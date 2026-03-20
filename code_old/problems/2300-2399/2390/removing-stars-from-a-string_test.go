package _2390

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_removeStars(t *testing.T) {
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
			if got := removeStars(tt.args.s); got != tt.want {
				t.Errorf("removeStars() = %v, want %v", got, tt.want)
			}
		})
	}
}
