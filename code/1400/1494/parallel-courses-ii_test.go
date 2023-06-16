package _1494

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_minNumberOfSemesters(t *testing.T) {
	type args struct {
		n         int
		relations [][]int
		k         int
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
			if got := minNumberOfSemesters(tt.args.n, tt.args.relations, tt.args.k); got != tt.want {
				t.Errorf("minNumberOfSemesters() = %v, want %v", got, tt.want)
			}
		})
	}
}
