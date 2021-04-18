package _0026

import (
	"github.com/godcong/leetcode/common"
	"testing"
)

func Test_removeDuplicates(t *testing.T) {
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
				nums: common.StrToIntArray("[1,1,2]"),
			},
			want: 2,
		},
		{
			name: "",
			args: args{
				nums: common.StrToIntArray("[0,0,1,1,1,2,2,3,3,4]"),
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := removeDuplicates(tt.args.nums); got != tt.want {
					t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
