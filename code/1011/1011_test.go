package _1011

import (
	"github.com/godcong/leetcode/common"
	"testing"
)

func Test_shipWithinDays(t *testing.T) {
	type args struct {
		weights []int
		D       int
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
				weights: common.StrToIntArray("[1,2,3,4,5,6,7,8,9,10]"),
				D:       5,
			},
			want: 15,
		},
		{
			name: "",
			args: args{
				weights: common.StrToIntArray("[3,2,2,4,1,4]"),
				D:       3,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shipWithinDays(tt.args.weights, tt.args.D); got != tt.want {
				t.Errorf("shipWithinDays() = %v, want %v", got, tt.want)
			}
		})
	}
}
