package _0081

import (
	. "github.com/godcong/leetcode/common"
	"testing"
)

func Test_search(t *testing.T) {
	type args struct {
		nums   []int
		target int
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
				nums:   StrToIntArray("[2,5,6,0,0,1,2]"),
				target: 0,
			},
			want: true,
		},
		{
			name: "",
			args: args{
				nums:   StrToIntArray("[2,5,6,0,0,1,2]"),
				target: 3,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
