package _2397

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_maximumRows(t *testing.T) {
	type args struct {
		matrix    [][]int
		numSelect int
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
			if got := maximumRows(tt.args.matrix, tt.args.numSelect); got != tt.want {
				t.Errorf("maximumRows() = %v, want %v", got, tt.want)
			}
		})
	}
}
