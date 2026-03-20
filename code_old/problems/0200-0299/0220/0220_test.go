package _0220

import (
	"github.com/godcong/leetcode/common"
	"testing"
)

func Test_containsNearbyAlmostDuplicate(t *testing.T) {
	type args struct {
		nums []int
		k    int
		t    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				nums: common.StrToIntArray("[1,2,3,1]"),
				k:    3,
				t:    0,
			},
			want: true,
		},
		{
			name: "",
			args: args{
				nums: common.StrToIntArray("[1,0,1,1]"),
				k:    1,
				t:    2,
			},
			want: true,
		},
		{
			name: "",
			args: args{
				nums: common.StrToIntArray("[1,5,9,1,5,9]"),
				k:    2,
				t:    3,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := containsNearbyAlmostDuplicate(tt.args.nums, tt.args.k, tt.args.t); got != tt.want {
					t.Errorf("containsNearbyAlmostDuplicate() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
