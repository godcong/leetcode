package _1504

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_numSubmat(t *testing.T) {
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
			if got := numSubmat(tt.args.mat); got != tt.want {
				t.Errorf("numSubmat() = %v, want %v", got, tt.want)
			}
		})
	}
}
