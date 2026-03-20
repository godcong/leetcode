package _1761

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_minTrioDegree(t *testing.T) {
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
			if got := minTrioDegree(tt.args.n, tt.args.edges); got != tt.want {
				t.Errorf("minTrioDegree() = %v, want %v", got, tt.want)
			}
		})
	}
}
