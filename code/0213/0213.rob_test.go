package _0213

import (
	"github.com/godcong/leetcode/common"
	"testing"
)

func Test_rob(t *testing.T) {
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
				nums: common.StrToIntArray("[2,3,2]"),
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := rob(tt.args.nums); got != tt.want {
					t.Errorf("rob() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
