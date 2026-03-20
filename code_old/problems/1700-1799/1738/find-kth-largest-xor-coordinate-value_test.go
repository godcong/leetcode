package _1738

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_kthLargestValue(t *testing.T) {
	type args struct {
		matrix [][]int
		k      int
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
			if got := kthLargestValue(tt.args.matrix, tt.args.k); got != tt.want {
				t.Errorf("kthLargestValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
