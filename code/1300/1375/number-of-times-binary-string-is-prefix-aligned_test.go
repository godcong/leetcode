package _1375

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_numTimesAllBlue(t *testing.T) {
	type args struct {
		flips []int
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
			if got := numTimesAllBlue(tt.args.flips); got != tt.want {
				t.Errorf("numTimesAllBlue() = %v, want %v", got, tt.want)
			}
		})
	}
}
