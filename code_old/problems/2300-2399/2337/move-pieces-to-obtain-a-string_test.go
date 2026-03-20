package _2337

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_canChange(t *testing.T) {
	type args struct {
		start  string
		target string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canChange(tt.args.start, tt.args.target); got != tt.want {
				t.Errorf("canChange() = %v, want %v", got, tt.want)
			}
		})
	}
}
