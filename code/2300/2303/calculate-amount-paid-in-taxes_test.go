package _2303

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_calculateTax(t *testing.T) {
	type args struct {
		brackets [][]int
		income   int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateTax(tt.args.brackets, tt.args.income); got != tt.want {
				t.Errorf("calculateTax() = %v, want %v", got, tt.want)
			}
		})
	}
}
