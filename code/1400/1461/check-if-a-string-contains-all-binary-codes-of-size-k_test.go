package _1461

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_hasAllCodes(t *testing.T) {
	type args struct {
		s string
		k int
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
			if got := hasAllCodes(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("hasAllCodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
