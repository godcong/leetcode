package _2427

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_commonFactors(t *testing.T) {
	type args struct {
		a int
		b int
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
			if got := commonFactors(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("commonFactors() = %v, want %v", got, tt.want)
			}
		})
	}
}
