package _1482

import (
	"github.com/godcong/leetcode/common"
	"testing"
)

func Test_minDays(t *testing.T) {
	type args struct {
		bloomDay []int
		m        int
		k        int
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
				bloomDay: common.StrToIntArray("[1,10,3,10,2]"),
				m:        3,
				k:        1,
			},
			want: 3,
		},
		{
			name: "",
			args: args{
				bloomDay: common.StrToIntArray("[1,10,3,10,2]"),
				m:        3,
				k:        2,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := minDays(tt.args.bloomDay, tt.args.m, tt.args.k); got != tt.want {
					t.Errorf("minDays() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
