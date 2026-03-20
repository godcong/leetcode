package _1582

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_numSpecial(t *testing.T) {
	type args struct {
		mat [][]int
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
			if got := numSpecial(tt.args.mat); got != tt.want {
				t.Errorf("numSpecial() = %v, want %v", got, tt.want)
			}
		})
	}
}
