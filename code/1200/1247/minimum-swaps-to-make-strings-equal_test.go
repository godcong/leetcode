package _1247

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_minimumSwap(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumSwap(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("minimumSwap() = %v, want %v", got, tt.want)
			}
		})
	}
}
