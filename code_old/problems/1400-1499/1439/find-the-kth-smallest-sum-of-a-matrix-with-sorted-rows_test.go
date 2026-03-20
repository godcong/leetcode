package _1439

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_kthSmallest(t *testing.T) {
	type args struct {
		mat [][]int
		k   int
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
			if got := kthSmallest(tt.args.mat, tt.args.k); got != tt.want {
				t.Errorf("kthSmallest() = %v, want %v", got, tt.want)
			}
		})
	}
}
