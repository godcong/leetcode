package _0877

import (
	"testing"
)

func Test_stoneGame(t *testing.T) {
	type args struct {
		piles []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stoneGame(tt.args.piles); got != tt.want {
				t.Errorf("stoneGame() = %v, want %v", got, tt.want)
			}
		})
	}
}
