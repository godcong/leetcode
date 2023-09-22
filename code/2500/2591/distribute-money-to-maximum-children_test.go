package _2591

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_distMoney(t *testing.T) {
	type args struct {
		money    int
		children int
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
			if got := distMoney(tt.args.money, tt.args.children); got != tt.want {
				t.Errorf("distMoney() = %v, want %v", got, tt.want)
			}
		})
	}
}
