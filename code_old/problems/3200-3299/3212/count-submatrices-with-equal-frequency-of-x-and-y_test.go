package _3212

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_numberOfSubmatrices(t *testing.T) {
	type args struct {
		grid [][]byte
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
			if got := numberOfSubmatrices(tt.args.grid); got != tt.want {
				t.Errorf("numberOfSubmatrices() = %v, want %v", got, tt.want)
			}
		})
	}
}
