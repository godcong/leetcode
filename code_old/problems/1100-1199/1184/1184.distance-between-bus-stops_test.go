package _1184

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_distanceBetweenBusStops(t *testing.T) {
	type args struct {
		distance    []int
		start       int
		destination int
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
			if got := distanceBetweenBusStops(tt.args.distance, tt.args.start, tt.args.destination); got != tt.want {
				t.Errorf("distanceBetweenBusStops() = %v, want %v", got, tt.want)
			}
		})
	}
}
