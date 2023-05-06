package _1419

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_minNumberOfFrogs(t *testing.T) {
	type args struct {
		croakOfFrogs string
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
			if got := minNumberOfFrogs(tt.args.croakOfFrogs); got != tt.want {
				t.Errorf("minNumberOfFrogs() = %v, want %v", got, tt.want)
			}
		})
	}
}
