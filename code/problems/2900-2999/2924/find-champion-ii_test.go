package _2924

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_findChampion(t *testing.T) {
	type args struct {
		n     int
		edges [][]int
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
			if got := findChampion(tt.args.n, tt.args.edges); got != tt.want {
				t.Errorf("findChampion() = %v, want %v", got, tt.want)
			}
		})
	}
}
