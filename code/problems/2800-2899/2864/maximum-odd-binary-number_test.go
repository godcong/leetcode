package _2864

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_maximumOddBinaryNumber(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumOddBinaryNumber(tt.args.s); got != tt.want {
				t.Errorf("maximumOddBinaryNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
