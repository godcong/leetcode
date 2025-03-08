package _2234

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_maximumBeauty(t *testing.T) {
	type args struct {
		flowers    []int
		newFlowers int64
		target     int
		full       int
		partial    int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumBeauty(tt.args.flowers, tt.args.newFlowers, tt.args.target, tt.args.full, tt.args.partial); got != tt.want {
				t.Errorf("maximumBeauty() = %v, want %v", got, tt.want)
			}
		})
	}
}
