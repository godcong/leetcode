package _1012

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_numDupDigitsAtMostN(t *testing.T) {
	type args struct {
		n int
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
			if got := numDupDigitsAtMostN(tt.args.n); got != tt.want {
				t.Errorf("numDupDigitsAtMostN() = %v, want %v", got, tt.want)
			}
		})
	}
}
