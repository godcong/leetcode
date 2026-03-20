package _2749

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_makeTheIntegerZero(t *testing.T) {
	type args struct {
		num1 int
		num2 int
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
			if got := makeTheIntegerZero(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("makeTheIntegerZero() = %v, want %v", got, tt.want)
			}
		})
	}
}
