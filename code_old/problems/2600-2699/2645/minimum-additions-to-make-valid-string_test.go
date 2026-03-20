package _2645

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_addMinimum(t *testing.T) {
	type args struct {
		word string
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
			if got := addMinimum(tt.args.word); got != tt.want {
				t.Errorf("addMinimum() = %v, want %v", got, tt.want)
			}
		})
	}
}
