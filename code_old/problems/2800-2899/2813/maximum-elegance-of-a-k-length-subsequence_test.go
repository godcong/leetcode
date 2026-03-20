package _2813

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_findMaximumElegance(t *testing.T) {
	type args struct {
		items [][]int
		k     int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaximumElegance(tt.args.items, tt.args.k); got != tt.want {
				t.Errorf("findMaximumElegance() = %v, want %v", got, tt.want)
			}
		})
	}
}
