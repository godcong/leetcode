package _3335

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_lengthAfterTransformations(t *testing.T) {
	type args struct {
		s string
		t int
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
			if got := lengthAfterTransformations(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("lengthAfterTransformations() = %v, want %v", got, tt.want)
			}
		})
	}
}
