package _3306

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_countOfSubstrings(t *testing.T) {
	type args struct {
		word string
		k    int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countOfSubstrings(tt.args.word, tt.args.k); got != tt.want {
				t.Errorf("countOfSubstrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
