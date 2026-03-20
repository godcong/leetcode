package _0926

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_minFlipsMonoIncr(t *testing.T) {
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
			if got := minFlipsMonoIncr(tt.args.s); got != tt.want {
				t.Errorf("minFlipsMonoIncr() = %v, want %v", got, tt.want)
			}
		})
	}
}
