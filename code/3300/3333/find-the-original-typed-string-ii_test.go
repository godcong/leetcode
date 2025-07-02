package _3333

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_possibleStringCount(t *testing.T) {
	type args struct {
		word string
		k    int
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
			if got := possibleStringCount(tt.args.word, tt.args.k); got != tt.want {
				t.Errorf("possibleStringCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
