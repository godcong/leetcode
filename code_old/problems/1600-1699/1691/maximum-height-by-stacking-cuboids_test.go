package _1691

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_maxHeight(t *testing.T) {
	type args struct {
		cuboids [][]int
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
			if got := maxHeight(tt.args.cuboids); got != tt.want {
				t.Errorf("maxHeight() = %v, want %v", got, tt.want)
			}
		})
	}
}
