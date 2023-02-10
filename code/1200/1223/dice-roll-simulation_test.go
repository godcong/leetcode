package _1223

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_dieSimulator(t *testing.T) {
	type args struct {
		n       int
		rollMax []int
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
			if got := dieSimulator(tt.args.n, tt.args.rollMax); got != tt.want {
				t.Errorf("dieSimulator() = %v, want %v", got, tt.want)
			}
		})
	}
}
