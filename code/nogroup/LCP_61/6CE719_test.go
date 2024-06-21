package _LCP_61

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_temperatureTrend(t *testing.T) {
	type args struct {
		temperatureA []int
		temperatureB []int
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
			if got := temperatureTrend(tt.args.temperatureA, tt.args.temperatureB); got != tt.want {
				t.Errorf("temperatureTrend() = %v, want %v", got, tt.want)
			}
		})
	}
}
