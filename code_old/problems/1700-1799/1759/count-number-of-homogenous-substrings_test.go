package _1759

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_countHomogenous(t *testing.T) {
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
			if got := countHomogenous(tt.args.s); got != tt.want {
				t.Errorf("countHomogenous() = %v, want %v", got, tt.want)
			}
		})
	}
}
