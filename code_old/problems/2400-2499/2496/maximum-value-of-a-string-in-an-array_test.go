package _2496

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_maximumValue(t *testing.T) {
	type args struct {
		strs []string
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
			if got := maximumValue(tt.args.strs); got != tt.want {
				t.Errorf("maximumValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
