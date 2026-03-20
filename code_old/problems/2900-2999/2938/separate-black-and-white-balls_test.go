package _2938

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_minimumSteps(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumSteps(tt.args.s); got != tt.want {
				t.Errorf("minimumSteps() = %v, want %v", got, tt.want)
			}
		})
	}
}
