package _2646

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_minimumTotalPrice(t *testing.T) {
	type args struct {
		n     int
		edges [][]int
		price []int
		trips [][]int
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
			if got := minimumTotalPrice(tt.args.n, tt.args.edges, tt.args.price, tt.args.trips); got != tt.want {
				t.Errorf("minimumTotalPrice() = %v, want %v", got, tt.want)
			}
		})
	}
}
