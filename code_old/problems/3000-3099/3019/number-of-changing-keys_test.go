package _3019

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_countKeyChanges(t *testing.T) {
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
			if got := countKeyChanges(tt.args.s); got != tt.want {
				t.Errorf("countKeyChanges() = %v, want %v", got, tt.want)
			}
		})
	}
}
