package _2006

import (
	"testing"
)

func Test_countKDifference(t *testing.T) {
	type args struct {
		nums []int
		k    int
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
			if got := countKDifference(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("countKDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
