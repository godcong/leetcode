package _1723

import (
	"github.com/godcong/leetcode/common"
	"testing"
)

func Test_minimumTimeRequired(t *testing.T) {
	type args struct {
		jobs []int
		k    int
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
				jobs: common.StrToIntArray("[3,2,3]"),
				k:    3,
			},
			want: 3,
		},
		{
			name: "",
			args: args{
				jobs: common.StrToIntArray("[1,2,4,7,8]"),
				k:    2,
			},
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := minimumTimeRequired(tt.args.jobs, tt.args.k); got != tt.want {
					t.Errorf("minimumTimeRequired() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
