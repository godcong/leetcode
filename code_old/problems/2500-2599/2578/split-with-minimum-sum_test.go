package _2578

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_splitNum(t *testing.T) {
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
			if got := splitNum(tt.args.num); got != tt.want {
				t.Errorf("splitNum() = %v, want %v", got, tt.want)
			}
		})
	}
}
