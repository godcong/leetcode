package _2806

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_accountBalanceAfterPurchase(t *testing.T) {
	type args struct {
		purchaseAmount int
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
			if got := accountBalanceAfterPurchase(tt.args.purchaseAmount); got != tt.want {
				t.Errorf("accountBalanceAfterPurchase() = %v, want %v", got, tt.want)
			}
		})
	}
}
