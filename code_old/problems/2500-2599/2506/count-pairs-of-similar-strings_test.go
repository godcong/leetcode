package _2506

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_similarPairs(t *testing.T) {
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
			if got := similarPairs(tt.args.words); got != tt.want {
				t.Errorf("similarPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
