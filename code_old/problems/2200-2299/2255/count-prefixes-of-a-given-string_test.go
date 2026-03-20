package _2255

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_countPrefixes(t *testing.T) {
	type args struct {
		words []string
		s     string
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
			if got := countPrefixes(tt.args.words, tt.args.s); got != tt.want {
				t.Errorf("countPrefixes() = %v, want %v", got, tt.want)
			}
		})
	}
}
