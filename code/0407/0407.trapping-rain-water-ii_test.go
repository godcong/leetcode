package _0407

import (
	"testing"
)

func Test_trapRainWater(t *testing.T) {
	type args struct {
		heightMap [][]int
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
			if got := trapRainWater(tt.args.heightMap); got != tt.want {
				t.Errorf("trapRainWater() = %v, want %v", got, tt.want)
			}
		})
	}
}
