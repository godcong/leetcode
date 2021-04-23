package _0368

import (
	"github.com/godcong/leetcode/common"
	"reflect"
	"testing"
)

func Test_largestDivisibleSubset(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				nums: common.StrToIntArray("[1,2,3]"),
			},
			want: common.StrToIntArray("[1,2]"),
		},
		{
			name: "",
			args: args{
				nums: common.StrToIntArray("[1,2,4,8]"),
			},
			want: common.StrToIntArray("[1,2,4,8]"),
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := largestDivisibleSubset(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("largestDivisibleSubset() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
