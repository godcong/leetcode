package _1690

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_stoneGameVII(t *testing.T) {
	type args struct {
		stones []int
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
			if got := stoneGameVII(tt.args.stones); got != tt.want {
				t.Errorf("stoneGameVII() = %v, want %v", got, tt.want)
			}
		})
	}
}
