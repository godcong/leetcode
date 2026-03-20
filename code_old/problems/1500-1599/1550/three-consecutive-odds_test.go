package _1550

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_threeConsecutiveOdds(t *testing.T) {
	type args struct {
		arr []int
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
			if got := threeConsecutiveOdds(tt.args.arr); got != tt.want {
				t.Errorf("threeConsecutiveOdds() = %v, want %v", got, tt.want)
			}
		})
	}
}
