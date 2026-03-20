package _3021

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_flowerGame(t *testing.T) {
	type args struct {
		n int
		m int
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
			if got := flowerGame(tt.args.n, tt.args.m); got != tt.want {
				t.Errorf("flowerGame() = %v, want %v", got, tt.want)
			}
		})
	}
}
