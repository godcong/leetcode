package _3233

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_nonSpecialCount(t *testing.T) {
	type args struct {
		l int
		r int
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
			if got := nonSpecialCount(tt.args.l, tt.args.r); got != tt.want {
				t.Errorf("nonSpecialCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
