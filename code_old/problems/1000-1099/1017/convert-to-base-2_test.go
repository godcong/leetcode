package _1017

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_baseNeg2(t *testing.T) {
	type args struct {
		n int
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
			if got := baseNeg2(tt.args.n); got != tt.want {
				t.Errorf("baseNeg2() = %v, want %v", got, tt.want)
			}
		})
	}
}
