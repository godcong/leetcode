package _1725

import (
	"testing"
)

func Test_countGoodRectangles(t *testing.T) {
	type args struct {
		rectangles [][]int
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
			if got := countGoodRectangles(tt.args.rectangles); got != tt.want {
				t.Errorf("countGoodRectangles() = %v, want %v", got, tt.want)
			}
		})
	}
}
