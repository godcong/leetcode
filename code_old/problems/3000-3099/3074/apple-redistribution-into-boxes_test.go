package _3074

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_minimumBoxes(t *testing.T) {
	type args struct {
		apple    []int
		capacity []int
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
			if got := minimumBoxes(tt.args.apple, tt.args.capacity); got != tt.want {
				t.Errorf("minimumBoxes() = %v, want %v", got, tt.want)
			}
		})
	}
}
