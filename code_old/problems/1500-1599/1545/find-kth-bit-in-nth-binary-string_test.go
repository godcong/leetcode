package _1545

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_findKthBit(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKthBit(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("findKthBit() = %v, want %v", got, tt.want)
			}
		})
	}
}
