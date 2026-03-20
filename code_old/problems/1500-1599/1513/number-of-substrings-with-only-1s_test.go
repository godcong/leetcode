package _1513

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_numSub(t *testing.T) {
	type args struct {
		s string
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
			if got := numSub(tt.args.s); got != tt.want {
				t.Errorf("numSub() = %v, want %v", got, tt.want)
			}
		})
	}
}
