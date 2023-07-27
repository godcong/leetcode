package _2500

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_deleteGreatestValue(t *testing.T) {
	type args struct {
		grid [][]int
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
			if got := deleteGreatestValue(tt.args.grid); got != tt.want {
				t.Errorf("deleteGreatestValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
