package _2451

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_oddString(t *testing.T) {
	type args struct {
		words []string
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
			if got := oddString(tt.args.words); got != tt.want {
				t.Errorf("oddString() = %v, want %v", got, tt.want)
			}
		})
	}
}
