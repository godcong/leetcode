package _2843

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_countSymmetricIntegers(t *testing.T) {
	type args struct {
		low  int
		high int
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
			if got := countSymmetricIntegers(tt.args.low, tt.args.high); got != tt.want {
				t.Errorf("countSymmetricIntegers() = %v, want %v", got, tt.want)
			}
		})
	}
}
