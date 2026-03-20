package _2423

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_equalFrequency(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := equalFrequency(tt.args.word); got != tt.want {
				t.Errorf("equalFrequency() = %v, want %v", got, tt.want)
			}
		})
	}
}
