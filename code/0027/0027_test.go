package _0027

import (
	"github.com/godcong/leetcode/common"
	"testing"
)

func Test_removeElement(t *testing.T) {
	type args struct {
		nums []int
		val  int
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
				nums: common.StrToIntArray("[3,2,2,3]"),
				val:  3,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := removeElement(tt.args.nums, tt.args.val); got != tt.want {
					t.Errorf("removeElement() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
