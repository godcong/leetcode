package _1995

import (
	"testing"
)

func Test_countQuadruplets(t *testing.T) {
	type args struct {
		nums []int
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
			if got := countQuadruplets(tt.args.nums); got != tt.want {
				t.Errorf("countQuadruplets() = %v, want %v", got, tt.want)
			}
		})
	}
}
