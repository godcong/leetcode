package _1742

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_countBalls(t *testing.T) {
	type args struct {
		lowLimit  int
		highLimit int
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
			if got := countBalls(tt.args.lowLimit, tt.args.highLimit); got != tt.want {
				t.Errorf("countBalls() = %v, want %v", got, tt.want)
			}
		})
	}
}
