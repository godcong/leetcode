package _0391

import (
	"testing"
)

func Test_isRectangleCover(t *testing.T) {
	type args struct {
		rectangles [][]int
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
			if got := isRectangleCover(tt.args.rectangles); got != tt.want {
				t.Errorf("isRectangleCover() = %v, want %v", got, tt.want)
			}
		})
	}
}
