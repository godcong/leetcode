package _0001

import (
	"github.com/godcong/leetcode/common"
	"reflect"
	"testing"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
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
				nums:   common.StrToIntArray("[2,7,11,15]"),
				target: 9,
			},
			want: common.StrToIntArray("[0,1]"),
		},
		{
			name: "",
			args: args{
				nums:   common.StrToIntArray("[3,2,4]"),
				target: 6,
			},
			want: common.StrToIntArray("[1,2]"),
		},
		{
			name: "",
			args: args{
				nums:   common.StrToIntArray("[3,3]"),
				target: 6,
			},
			want: common.StrToIntArray("[0,1]"),
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := twoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("twoSum() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
