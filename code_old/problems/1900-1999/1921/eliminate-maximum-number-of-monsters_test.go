package _1921

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_eliminateMaximum(t *testing.T) {
	type args struct {
		dist  []int
		speed []int
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
			if got := eliminateMaximum(tt.args.dist, tt.args.speed); got != tt.want {
				t.Errorf("eliminateMaximum() = %v, want %v", got, tt.want)
			}
		})
	}
}
