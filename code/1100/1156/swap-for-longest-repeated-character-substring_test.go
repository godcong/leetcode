package _1156

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_maxRepOpt1(t *testing.T) {
	type args struct {
		text string
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
			if got := maxRepOpt1(tt.args.text); got != tt.want {
				t.Errorf("maxRepOpt1() = %v, want %v", got, tt.want)
			}
		})
	}
}
