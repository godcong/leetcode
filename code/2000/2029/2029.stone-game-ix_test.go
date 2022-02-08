package _2029

import (
	"testing"
)

func Test_stoneGameIX(t *testing.T) {
	type args struct {
		stones []int
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
			if got := stoneGameIX(tt.args.stones); got != tt.want {
				t.Errorf("stoneGameIX() = %v, want %v", got, tt.want)
			}
		})
	}
}
