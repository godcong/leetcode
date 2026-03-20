package _3494

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_minTime(t *testing.T) {
	type args struct {
		skill []int
		mana  []int
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
			if got := minTime(tt.args.skill, tt.args.mana); got != tt.want {
				t.Errorf("minTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
