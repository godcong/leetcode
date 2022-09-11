package _0857

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_mincostToHireWorkers(t *testing.T) {
	type args struct {
		quality []int
		wage    []int
		k       int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mincostToHireWorkers(tt.args.quality, tt.args.wage, tt.args.k); got != tt.want {
				t.Errorf("mincostToHireWorkers() = %v, want %v", got, tt.want)
			}
		})
	}
}
