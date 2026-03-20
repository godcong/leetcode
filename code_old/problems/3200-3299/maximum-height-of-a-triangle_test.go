package _3200

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_maxHeightOfTriangle(t *testing.T) {
	type args struct {
		red  int
		blue int
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
			if got := maxHeightOfTriangle(tt.args.red, tt.args.blue); got != tt.want {
				t.Errorf("maxHeightOfTriangle() = %v, want %v", got, tt.want)
			}
		})
	}
}
