package _2383

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_minNumberOfHours(t *testing.T) {
	type args struct {
		initialEnergy     int
		initialExperience int
		energy            []int
		experience        []int
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
			if got := minNumberOfHours(tt.args.initialEnergy, tt.args.initialExperience, tt.args.energy, tt.args.experience); got != tt.want {
				t.Errorf("minNumberOfHours() = %v, want %v", got, tt.want)
			}
		})
	}
}
