package _1630

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_checkArithmeticSubarrays(t *testing.T) {
	type args struct {
		nums []int
		l    []int
		r    []int
	}
	tests := []struct {
		name string
		args args
		want []bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkArithmeticSubarrays(tt.args.nums, tt.args.l, tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("checkArithmeticSubarrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
