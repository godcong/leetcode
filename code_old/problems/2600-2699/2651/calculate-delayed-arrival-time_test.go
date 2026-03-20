package _2651

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_findDelayedArrivalTime(t *testing.T) {
	type args struct {
		arrivalTime int
		delayedTime int
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
			if got := findDelayedArrivalTime(tt.args.arrivalTime, tt.args.delayedTime); got != tt.want {
				t.Errorf("findDelayedArrivalTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
