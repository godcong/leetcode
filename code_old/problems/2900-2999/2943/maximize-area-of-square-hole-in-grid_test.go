package _2943

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_maximizeSquareHoleArea(t *testing.T) {
	type args struct {
		n     int
		m     int
		hBars []int
		vBars []int
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
			if got := maximizeSquareHoleArea(tt.args.n, tt.args.m, tt.args.hBars, tt.args.vBars); got != tt.want {
				t.Errorf("maximizeSquareHoleArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
