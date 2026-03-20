package _2335

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_fillCups(t *testing.T) {
	type args struct {
		amount []int
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
			if got := fillCups(tt.args.amount); got != tt.want {
				t.Errorf("fillCups() = %v, want %v", got, tt.want)
			}
		})
	}
}
