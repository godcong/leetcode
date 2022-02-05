package _1219

import (
	"testing"
)

func Test_getMaximumGold(t *testing.T) {
	type args struct {
		grid [][]int
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
			if got := getMaximumGold(tt.args.grid); got != tt.want {
				t.Errorf("getMaximumGold() = %v, want %v", got, tt.want)
			}
		})
	}
}
