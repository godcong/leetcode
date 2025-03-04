package _1745

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_checkPartitioning(t *testing.T) {
	type args struct {
		s string
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
			if got := checkPartitioning(tt.args.s); got != tt.want {
				t.Errorf("checkPartitioning() = %v, want %v", got, tt.want)
			}
		})
	}
}
