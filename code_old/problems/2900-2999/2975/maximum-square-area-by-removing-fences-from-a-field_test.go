package _2975

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_maximizeSquareArea(t *testing.T) {
	type args struct {
		m       int
		n       int
		hFences []int
		vFences []int
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
			if got := maximizeSquareArea(tt.args.m, tt.args.n, tt.args.hFences, tt.args.vFences); got != tt.want {
				t.Errorf("maximizeSquareArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
