package _LCP_33

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_storeWater(t *testing.T) {
	type args struct {
		bucket []int
		vat    []int
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
			if got := storeWater(tt.args.bucket, tt.args.vat); got != tt.want {
				t.Errorf("storeWater() = %v, want %v", got, tt.want)
			}
		})
	}
}
