package _1806

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_reinitializePermutation(t *testing.T) {
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
			if got := reinitializePermutation(tt.args.n); got != tt.want {
				t.Errorf("reinitializePermutation() = %v, want %v", got, tt.want)
			}
		})
	}
}
