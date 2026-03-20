package _1969

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_minNonZeroProduct(t *testing.T) {
	type args struct {
		p int
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
			if got := minNonZeroProduct(tt.args.p); got != tt.want {
				t.Errorf("minNonZeroProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
