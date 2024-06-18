package _2288

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_discountPrices(t *testing.T) {
	type args struct {
		sentence string
		discount int
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
			if got := discountPrices(tt.args.sentence, tt.args.discount); got != tt.want {
				t.Errorf("discountPrices() = %v, want %v", got, tt.want)
			}
		})
	}
}
