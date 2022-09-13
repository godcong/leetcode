package _0670

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_maximumSwap(t *testing.T) {
	type args struct {
		num int
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
			if got := maximumSwap(tt.args.num); got != tt.want {
				t.Errorf("maximumSwap() = %v, want %v", got, tt.want)
			}
		})
	}
}
