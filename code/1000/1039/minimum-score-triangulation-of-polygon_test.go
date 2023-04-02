package _1039

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_minScoreTriangulation(t *testing.T) {
	type args struct {
		values []int
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
			if got := minScoreTriangulation(tt.args.values); got != tt.want {
				t.Errorf("minScoreTriangulation() = %v, want %v", got, tt.want)
			}
		})
	}
}
