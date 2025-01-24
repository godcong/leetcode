package _2412

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_minimumMoney(t *testing.T) {
	type args struct {
		transactions [][]int
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
			if got := minimumMoney(tt.args.transactions); got != tt.want {
				t.Errorf("minimumMoney() = %v, want %v", got, tt.want)
			}
		})
	}
}
