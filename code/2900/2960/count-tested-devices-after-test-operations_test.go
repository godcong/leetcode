package _2960

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_countTestedDevices(t *testing.T) {
	type args struct {
		batteryPercentages []int
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
			if got := countTestedDevices(tt.args.batteryPercentages); got != tt.want {
				t.Errorf("countTestedDevices() = %v, want %v", got, tt.want)
			}
		})
	}
}
