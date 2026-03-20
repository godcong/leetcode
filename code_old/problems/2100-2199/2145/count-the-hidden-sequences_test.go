package _2145

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_numberOfArrays(t *testing.T) {
	type args struct {
		differences []int
		lower       int
		upper       int
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
			if got := numberOfArrays(tt.args.differences, tt.args.lower, tt.args.upper); got != tt.want {
				t.Errorf("numberOfArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
