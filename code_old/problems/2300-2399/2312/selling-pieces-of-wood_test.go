package _2312

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_sellingWood(t *testing.T) {
	type args struct {
		m      int
		n      int
		prices [][]int
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
			if got := sellingWood(tt.args.m, tt.args.n, tt.args.prices); got != tt.want {
				t.Errorf("sellingWood() = %v, want %v", got, tt.want)
			}
		})
	}
}
