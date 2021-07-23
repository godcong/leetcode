package _1893

import (
	"testing"
)

func Test_isCovered(t *testing.T) {
	type args struct {
		ranges [][]int
		left   int
		right  int
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
			if got := isCovered(tt.args.ranges, tt.args.left, tt.args.right); got != tt.want {
				t.Errorf("isCovered() = %v, want %v", got, tt.want)
			}
		})
	}
}
