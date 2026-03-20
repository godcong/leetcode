package _1801

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_getNumberOfBacklogOrders(t *testing.T) {
	type args struct {
		orders [][]int
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
			if got := getNumberOfBacklogOrders(tt.args.orders); got != tt.want {
				t.Errorf("getNumberOfBacklogOrders() = %v, want %v", got, tt.want)
			}
		})
	}
}
