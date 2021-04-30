package _0137

import (
	"github.com/godcong/leetcode/common"
	"testing"
)

func Test_singleNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				nums: common.StrToIntArray("[2,2,3,2]"),
			},
			want: 3,
		},
		{
			name: "",
			args: args{
				nums: common.StrToIntArray("[0,1,0,1,0,1,99]"),
			},
			want: 99,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := singleNumber(tt.args.nums); got != tt.want {
					t.Errorf("singleNumber() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
