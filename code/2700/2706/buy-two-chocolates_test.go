package _2706

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_buyChoco(t *testing.T) {
	type args struct {
		prices []int
		money  int
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
			if got := buyChoco(tt.args.prices, tt.args.money); got != tt.want {
				t.Errorf("buyChoco() = %v, want %v", got, tt.want)
			}
		})
	}
}
