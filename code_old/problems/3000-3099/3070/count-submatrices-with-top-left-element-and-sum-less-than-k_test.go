package _3070

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_countSubmatrices(t *testing.T) {
	type args struct {
		grid [][]int
		k    int
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
			if got := countSubmatrices(tt.args.grid, tt.args.k); got != tt.want {
				t.Errorf("countSubmatrices() = %v, want %v", got, tt.want)
			}
		})
	}
}
