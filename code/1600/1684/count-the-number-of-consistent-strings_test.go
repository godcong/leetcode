package _1684

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_countConsistentStrings(t *testing.T) {
	type args struct {
		allowed string
		words   []string
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
			if got := countConsistentStrings(tt.args.allowed, tt.args.words); got != tt.want {
				t.Errorf("countConsistentStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
