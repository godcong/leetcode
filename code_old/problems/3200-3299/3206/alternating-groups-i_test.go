package _3206

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_numberOfAlternatingGroups(t *testing.T) {
	type args struct {
		colors []int
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
			if got := numberOfAlternatingGroups(tt.args.colors); got != tt.want {
				t.Errorf("numberOfAlternatingGroups() = %v, want %v", got, tt.want)
			}
		})
	}
}
