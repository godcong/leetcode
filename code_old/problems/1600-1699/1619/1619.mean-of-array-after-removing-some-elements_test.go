package _1619

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_trimMean(t *testing.T) {
	type args struct {
		arr []int
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
			if got := trimMean(tt.args.arr); got != tt.want {
				t.Errorf("trimMean() = %v, want %v", got, tt.want)
			}
		})
	}
}
