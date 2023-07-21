package _1499

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_findMaxValueOfEquation(t *testing.T) {
	type args struct {
		points [][]int
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
			if got := findMaxValueOfEquation(tt.args.points, tt.args.k); got != tt.want {
				t.Errorf("findMaxValueOfEquation() = %v, want %v", got, tt.want)
			}
		})
	}
}
