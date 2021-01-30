package code

import "testing"

func Test_swimInWater(t *testing.T) {
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
			if got := swimInWater(tt.args.grid); got != tt.want {
				t.Errorf("swimInWater() = %v, want %v", got, tt.want)
			}
		})
	}
}
