package _2739

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_distanceTraveled(t *testing.T) {
	type args struct {
		mainTank       int
		additionalTank int
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
			if got := distanceTraveled(tt.args.mainTank, tt.args.additionalTank); got != tt.want {
				t.Errorf("distanceTraveled() = %v, want %v", got, tt.want)
			}
		})
	}
}
