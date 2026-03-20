package _2376

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_countSpecialNumbers(t *testing.T) {
	type args struct {
		n int
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
			if got := countSpecialNumbers(tt.args.n); got != tt.want {
				t.Errorf("countSpecialNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
