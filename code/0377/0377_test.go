package _0377

import (
	"github.com/godcong/leetcode/common"
	"testing"
)

func Test_combinationSum4(t *testing.T) {
	type args struct {
		nums   []int
		target int
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
				nums:   common.StrToIntArray("[1,2,3]"),
				target: 6,
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := combinationSum4(tt.args.nums, tt.args.target); got != tt.want {
					t.Errorf("combinationSum4() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
