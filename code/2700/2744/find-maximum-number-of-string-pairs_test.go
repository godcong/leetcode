package _2744

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_maximumNumberOfStringPairs(t *testing.T) {
	type args struct {
		words []string
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
			if got := maximumNumberOfStringPairs(tt.args.words); got != tt.want {
				t.Errorf("maximumNumberOfStringPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
