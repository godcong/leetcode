package _3194

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_minimumAverage(t *testing.T) {
	type args struct {
		nums []int
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
			if got := minimumAverage(tt.args.nums); got != tt.want {
				t.Errorf("minimumAverage() = %v, want %v", got, tt.want)
			}
		})
	}
}
