package _0879

import (
	"testing"
)

func Test_profitableSchemes(t *testing.T) {
	type args struct {
		n         int
		minProfit int
		group     []int
		profit    []int
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
			if got := profitableSchemes(tt.args.n, tt.args.minProfit, tt.args.group, tt.args.profit); got != tt.want {
				t.Errorf("profitableSchemes() = %v, want %v", got, tt.want)
			}
		})
	}
}
