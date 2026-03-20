package _3099

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_sumOfTheDigitsOfHarshadNumber(t *testing.T) {
	type args struct {
		x int
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
			if got := sumOfTheDigitsOfHarshadNumber(tt.args.x); got != tt.want {
				t.Errorf("sumOfTheDigitsOfHarshadNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
