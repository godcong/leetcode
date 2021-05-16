package _0421

import (
	"testing"

	"github.com/godcong/leetcode/common"
)

func Test_findMaximumXOR(t *testing.T) {
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
				nums: common.StrToIntArray("[3,10,5,25,2,8]"),
			},
			want: 28,
		},
		{
			name: "",
			args: args{
				nums: common.StrToIntArray("[0]"),
			},
			want: 0,
		},
		{
			name: "",
			args: args{
				nums: common.StrToIntArray("[2,4]"),
			},
			want: 6,
		},
		{
			name: "",
			args: args{
				nums: common.StrToIntArray("[8,10,2]"),
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaximumXOR(tt.args.nums); got != tt.want {
				t.Errorf("findMaximumXOR() = %v, want %v", got, tt.want)
			}
		})
	}
}
