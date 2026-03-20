package _2432

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_hardestWorker(t *testing.T) {
	type args struct {
		n    int
		logs [][]int
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
			if got := hardestWorker(tt.args.n, tt.args.logs); got != tt.want {
				t.Errorf("hardestWorker() = %v, want %v", got, tt.want)
			}
		})
	}
}
