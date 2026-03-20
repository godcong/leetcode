package _1687

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_boxDelivering(t *testing.T) {
	type args struct {
		boxes      [][]int
		portsCount int
		maxBoxes   int
		maxWeight  int
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
			if got := boxDelivering(tt.args.boxes, tt.args.portsCount, tt.args.maxBoxes, tt.args.maxWeight); got != tt.want {
				t.Errorf("boxDelivering() = %v, want %v", got, tt.want)
			}
		})
	}
}
