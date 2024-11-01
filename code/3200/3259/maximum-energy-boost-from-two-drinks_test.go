package _3259

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_maxEnergyBoost(t *testing.T) {
	type args struct {
		energyDrinkA []int
		energyDrinkB []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxEnergyBoost(tt.args.energyDrinkA, tt.args.energyDrinkB); got != tt.want {
				t.Errorf("maxEnergyBoost() = %v, want %v", got, tt.want)
			}
		})
	}
}
