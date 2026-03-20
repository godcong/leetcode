package _1870

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_minSpeedOnTime(t *testing.T) {
	type args struct {
		dist []int
		hour float64
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
			if got := minSpeedOnTime(tt.args.dist, tt.args.hour); got != tt.want {
				t.Errorf("minSpeedOnTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
